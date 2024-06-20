package dates

import (
	"reflect"
	"testing"
)

func TestFindDateRanges(t *testing.T) {
	tests := []struct {
		name      string
		startDate string
		endDate   string
		dayBegin  Weekday
		dayStart  Weekday
		want      []DateRange
		wantErr   bool
	}{
		{
			name:      "Wednesday to Tuesday range",
			startDate: "2023-06-01",
			endDate:   "2023-06-30",
			dayBegin:  Wednesday,
			dayStart:  Tuesday,
			want: []DateRange{
				{BeginDate: "2023-06-07", EndDate: "2023-06-13"},
				{BeginDate: "2023-06-14", EndDate: "2023-06-20"},
				{BeginDate: "2023-06-21", EndDate: "2023-06-27"},
				{BeginDate: "2023-06-28", EndDate: "2023-06-30"},
			},
			wantErr: false,
		},
		{
			name:      "Friday to Thursday range",
			startDate: "2023-07-01",
			endDate:   "2023-07-31",
			dayBegin:  Friday,
			dayStart:  Thursday,
			want: []DateRange{
				{BeginDate: "2023-07-07", EndDate: "2023-07-13"},
				{BeginDate: "2023-07-14", EndDate: "2023-07-20"},
				{BeginDate: "2023-07-21", EndDate: "2023-07-27"},
				{BeginDate: "2023-07-28", EndDate: "2023-07-31"},
			},
			wantErr: false,
		},
		{
			name:      "Start date is dayBegin (Tuesday)",
			startDate: "2023-08-01", // Tuesday
			endDate:   "2023-08-14",
			dayBegin:  Tuesday,
			dayStart:  Monday,
			want: []DateRange{
				{BeginDate: "2023-08-01", EndDate: "2023-08-07"},
				{BeginDate: "2023-08-08", EndDate: "2023-08-14"},
			},
			wantErr: false,
		},
		{
			name:      "End date is dayStart (Saturday)",
			startDate: "2023-09-01",
			endDate:   "2023-09-16", // Saturday
			dayBegin:  Sunday,
			dayStart:  Saturday,
			want: []DateRange{
				{BeginDate: "2023-09-03", EndDate: "2023-09-09"},
				{BeginDate: "2023-09-10", EndDate: "2023-09-16"},
			},
			wantErr: false,
		},
		{
			name:      "Single day range (Thursday)",
			startDate: "2023-10-05", // Thursday
			endDate:   "2023-10-05", // Thursday
			dayBegin:  Thursday,
			dayStart:  Friday,
			want: []DateRange{
				{BeginDate: "2023-10-05", EndDate: "2023-10-05"},
			},
			wantErr: false,
		},
		{
			name:      "Start date after end date",
			startDate: "2023-11-30",
			endDate:   "2023-11-01",
			dayBegin:  Wednesday,
			dayStart:  Tuesday,
			want:      []DateRange{},
			wantErr:   false,
		},
		{
			name:      "Invalid start date format",
			startDate: "2023/12/01",
			endDate:   "2023-12-31",
			dayBegin:  Friday,
			dayStart:  Thursday,
			want:      nil,
			wantErr:   true,
		},
		{
			name:      "Invalid end date format",
			startDate: "2024-01-01",
			endDate:   "2024/01/31",
			dayBegin:  Monday,
			dayStart:  Sunday,
			want:      nil,
			wantErr:   true,
		},
		{
			name:      "Leap year case (Saturday to Friday)",
			startDate: "2024-02-24", // Saturday
			endDate:   "2024-03-08", // Friday
			dayBegin:  Saturday,
			dayStart:  Friday,
			want: []DateRange{
				{BeginDate: "2024-02-24", EndDate: "2024-03-01"},
				{BeginDate: "2024-03-02", EndDate: "2024-03-08"},
			},
			wantErr: false,
		},
		{
			name:      "Year boundary case (Thursday to Wednesday)",
			startDate: "2023-12-28", // Thursday
			endDate:   "2024-01-10", // Wednesday
			dayBegin:  Thursday,
			dayStart:  Wednesday,
			want: []DateRange{
				{BeginDate: "2023-12-28", EndDate: "2024-01-03"},
				{BeginDate: "2024-01-04", EndDate: "2024-01-10"},
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := FindDateRanges(tt.startDate, tt.endDate, tt.dayBegin, tt.dayStart)
			if (err != nil) != tt.wantErr {
				t.Errorf("FindDateRanges() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FindDateRanges() = %v, want %v", got, tt.want)
			}
		})
	}
}
