CREATE TABLE TZodiak (
    ID SERIAL PRIMARY KEY,
    StartDate DATE,
    EndDate DATE,
    ZodiacName VARCHAR(50)
);


INSERT INTO TZodiak (StartDate, EndDate, ZodiacName)
VALUES
    ('2023-12-22', '2024-01-19', 'Capricorn'),
    ('2024-01-20', '2024-02-18', 'Aquarius'),
    ('2024-02-19', '2024-03-20', 'Pisces'),
    ('2024-03-21', '2024-04-19', 'Aries'),
    ('2024-04-20', '2024-05-20', 'Taurus'),
    ('2024-05-21', '2024-06-20', 'Gemini'),
    ('2024-06-21', '2024-07-22', 'Cancer'),
    ('2024-07-23', '2024-08-22', 'Leo'),
    ('2024-08-23', '2024-09-22', 'Virgo'),
    ('2024-09-23', '2024-10-22', 'Libra'),
    ('2024-10-23', '2024-11-21', 'Scorpio'),
    ('2024-11-22', '2024-12-21', 'Sagittarius');
