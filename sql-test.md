**[Hard]**
Show patient_id, first_name, last_name, and attending doctor's specialty.
Show only the patients who has a diagnosis as 'Epilepsy' and the doctor's first name is 'Lisa'
Check patients, admissions, and doctors tables for required information.

```sql
SELECT
    p.patient_id AS patient_id,
    p.first_name AS patient_first_name,
    p.last_name AS patient_last_name,
    d.specialty AS attending_doctor_special
FROM patients AS p
INNER JOIN admissions as a ON a.diagnosis = "Epilepsy" AND a.patient_id = p.patient_id
INNER JOIN doctors AS d ON d.first_name = "Lisa" AND d.doctor_id = a.attending_doctor_id
```
---
**[Hard]**
Show all of the patients grouped into weight groups.
Show the total amount of patients in each weight group.
Order the list by the weight group decending.

For example, if they weight 100 to 109 they are placed in the 100 weight group, 110-119 = 110 weight group, etc.

```sql
SELECT
    COUNT(),
    weight / 10 * 10 AS weight_group
FROM patients
GROUP BY weight_group
ORDER BY weight_group DESC
```
---
**[Medium]**
Show patient_id and first_name from patients where their first_name start and ends with 's' and is at least 6 characters long.

```sql
SELECT
    patient_id,
    first_name
FROM patients
WHERE
    first_name LIKE "s%s" AND
    LEN(first_name) >= 6
```
---
**[Medium]**
Show unique birth years from patients and order them by ascending.

```sql
SELECT
    DISTINCT(CAST(birth_date AS YEAR)) AS birth_date
FROM patients
ORDER BY birth_date ASC
```
---
**[Medium]**
Show unique first names from the patients table which only occurs once in the list.
For example, if two or more people are named 'John' in the first_name column then don't include their name in the output list. If only 1 person is named 'Leo' then include them in the output.

```sql
SELECT
    first_name
FROM patients
GROUP BY first_name
HAVING COUNT(first_name) = 1
```
---