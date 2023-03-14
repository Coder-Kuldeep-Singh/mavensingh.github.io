#include <stdio.h>

int main()
{
    int year, month, day;
    int daysInMonth, startingDay;

    printf("Enter the year: ");
    scanf("%d", &year);

    char *months[] = {"January", "February", "March", "April", "May", "June", "July", "August", "September", "October", "November", "December"};

    int daysInWeek = 7;
    int calendar[6][7];

    // Calculate starting day
    startingDay = getStartingDay(year);

    // Fill in the calendar with dates
    int row = 0;
    for (month = 0; month < 12; month++)
    {
        daysInMonth = getDaysInMonth(month, year);

        // Print month and year header
        printf("\n     %s %d\n", months[month], year);
        printf("Sun Mon Tue Wed Thu Fri Sat\n");

        // Fill in the days for the month
        for (day = 1; day <= daysInMonth; day++)
        {
            calendar[row][startingDay] = day;
            startingDay++;

            // Move to the next row if necessary
            if (startingDay > 6)
            {
                startingDay = 0;
                row++;
            }
        }

        // Print the calendar for the month
        for (int i = 0; i < 6; i++)
        {
            for (int j = 0; j < 7; j++)
            {
                if (calendar[i][j] == 0)
                {
                    printf("    ");
                }
                else
                {
                    printf("%3d ", calendar[i][j]);
                }
            }
            printf("\n");
        }

        // Reset the row and starting day for the next month
        row = 0;
        startingDay = (startingDay + daysInMonth) % daysInWeek;
    }
}

int getStartingDay(int year)
{
    int d = (1 + (year - 1) * 365 + (year - 1) / 4 - (year - 1) / 100 + (year - 1) / 400) % 7;
    return d;
}

int getDaysInMonth(int month, int year)
{
    int daysInMonth;

    if (month == 1)
    {
        if ((year % 4 == 0 && year % 100 != 0) || year % 400 == 0)
        {
            daysInMonth = 29;
        }
        else
        {
            daysInMonth = 28;
        }
    }
    else if (month == 3 || month == 5 || month == 8 || month == 10)
    {
        daysInMonth = 30;
    }
    else
    {
        daysInMonth = 31;
    }

    return daysInMonth;
}
