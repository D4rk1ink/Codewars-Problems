**Problem**

ให้หา sum ของ 2 string value โดยที่ string value จะเป็น big integer ซึ่งไม่สามารถใช้ operation + ที่มีข้อกำจัดของขนาด integer ได้

```
Expected: '91002328220491911630239667963', instead got: '9.100232822049192e+28'
```

**Solution**

จะใช้วิธีไล่บวกแต่ละหลักตามหลักคณิตศาสตร์ตามหลักปกติ

โดย variable แต่ละตัวจะทำหน้าที่ตามหลักคณิตศาสตร์ดังนี้
```go
result := "" // ใช้เก็บผลลัพธ์ที่เกิดจากการคำนวนในแต่ละหลัก
pivot := 1 // ใช้เก็บหลักที่กำหนดคำนวน
extended := 0 // ใช้เก็บเลขทดที่เกิดจากการคำนวน
```

```go
if pivot <= lenNum1 {
    n1, _ = strconv.Atoi(string(num1[lenNum1-pivot]))
}
if pivot <= lenNum2 {
    n2, _ = strconv.Atoi(string(num2[lenNum2-pivot]))
}
```
จาก code จะทำการดึงแต่ละหลักมาคำนวน เริ่มจากหลักหน่วย สิบ ร้อย ไปเรื่อยตาม pivot

```go
cal := n1 + n2 + extended
```
ทำการบวกค่าที่ได้จากแต่ละหลัก โดยถ้ามีแทนทดที่ carry มาด้วยก็จะนำมาคำนวนด้วย

```go
result = strconv.Itoa(cal%10) + result
extended = int(math.Floor(float64(cal / 10)))

pivot += 1
```
กรณีบวกและได้ค่าเกินสิบ จะทำการนำเลขหลักหน่วยมาเติมในผลลัพธ์ และทำการเก็บเลขทดไว้ในตัวแปร `extended` และวนทำแบบนี้ไปทุก ๆ หลัก

```go
if extended != 0 {
    result = strconv.Itoa(extended) + result
}
```

เมื่อทำการวนคำนวนทุกหลักแล้วและเหลือเลขทด จะนำเลขทดมาเติมในข้างหน้าผลลัพธ์