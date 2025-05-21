**Codewars Problem: Multiplying numbers as strings**

https://www.codewars.com/kata/55911ef14065454c75000062

**Codewars Result**

![result](https://github.com/D4rk1ink/Codewars-Problems/blob/master/multiplying-numbers-as-strings/result.png?raw=true)

**Problem**

ให้หา result ของการคูณของทั้ว 2 string value โดยที่ string value จะเป็น big integer ซึ่งไม่สามารถใช้ operation * ที่มีข้อกำจัดของขนาด integer ได้

```
Expected: '91002328220491911630239667963', instead got: '9.100232822049192e+28'
```

**Solution**

จะใช้วิธีไล่คูณแต่ละหลักตามหลักคณิตศาสตร์ตามหลักปกติ

เริ่มตันจำการเช็คว่า multiplicand หรือ multiplier เป็น 0 ไหม ถ้าเป็น 0 จะ return 0 กลับไปเลยเพื่อลด process
```go
if multiplicand == "0" || multiplier == "0" {
    return "0"
}
```

โดย code หลักจะเป็นการนำแต่ละหลักของ multiplier (ตัวคูณ) มาไล่คูณกับแต่ละหลักของ multiplicand (ตัวตั้ง)

```go
for i := lenMultiplier - 1; i >= 0; i-- {
    m1, _ := strconv.Atoi(string(multiplier[i]))

    for j := lenMultiplicand - 1; j >= 0; j-- {
        m2, _ := strconv.Atoi(string(multiplicand[j]))
        cal := m1*m2 + extended
        
        ...
    }

    ...
}
```

ผลลัพธ์ที่ถูกคำนวนในแต่ละหลักของ multiplier จะเริ่มจากหลักของ multiplier ที่กำลังคำนวนอยู่ เช่น กำลังคำนวนอยู่ที่หลักสิบ ผลลัพธ์เริ่มต้นก็จะอยู่ที่ 10 (หรือคือการ reserve 0 ทิ้งไว้ก่อน)


```diff
for i := lenMultiplier - 1; i >= 0; i-- {
    m1, _ := strconv.Atoi(string(multiplier[i]))

+    for z := 0; z < (lenMultiplier - 1 - i); z++ {
+        result += "0"
+    }

    for j := lenMultiplicand - 1; j >= 0; j-- {
        m2, _ := strconv.Atoi(string(multiplicand[j]))
        cal := m1*m2 + extended
        
        ...
    }

    ...
}
```

code จะทำการดึงแต่ละหลักมาคำนวน เริ่มจากหลักหน่วย สิบ ร้อย ไปเรื่อยตาม pivot

ทำการคูณค่าที่ได้จากแต่ละหลัก โดยถ้ามีแทนทดที่ carry มาด้วยก็จะนำมาคำนวนด้วย

กรณีคูณและได้ค่าเกินสิบ จะทำการนำเลขหลักหน่วยมาเติมในผลลัพธ์ และทำการเก็บเลขทดไว้ในตัวแปร `extended` และวนทำแบบนี้ไปทุก ๆ หลัก

```diff
for i := lenMultiplier - 1; i >= 0; i-- {
    m1, _ := strconv.Atoi(string(multiplier[i]))

    for z := 0; z < (lenMultiplier - 1 - i); z++ {
        result += "0"
    }

    for j := lenMultiplicand - 1; j >= 0; j-- {
        m2, _ := strconv.Atoi(string(multiplicand[j]))
        cal := m1*m2 + extended

+       if m2 == 0 && extended != 0 {
+           result = strconv.Itoa(cal) + result
+           extended = 0
+           lastNonZeroIndex = j
+       } else {
+           result = strconv.Itoa(cal%10) + result
+           extended = int(math.Floor(float64(cal / 10)))
+       }
    }

    ...
}
```

หลังนำ multiplier แต่ละหลักคูณกับ multiplicand แต่ละหลักหมดแล้วและเหลือเลขทด จะนำเลขทดมาเติมในข้างหน้าผลลัพธ์ ในแต่ละผลลัพธ์ของตัวคูณ และนำผลลัพธ์นั้นบวกเพิ่มไปยังผลลัพธ์ก่อนหน้า ทำแบบนี้ไปเรื่อย ๆ จน multiplier และ multiplicand คูณกันหมด

```diff
for i := lenMultiplier - 1; i >= 0; i-- {
    m1, _ := strconv.Atoi(string(multiplier[i]))

    for z := 0; z < (lenMultiplier - 1 - i); z++ {
        result += "0"
    }

    for j := lenMultiplicand - 1; j >= 0; j-- {
        m2, _ := strconv.Atoi(string(multiplicand[j]))
        cal := m1*m2 + extended

        if m2 == 0 && extended != 0 {
            result = strconv.Itoa(cal) + result
            extended = 0
            lastNonZeroIndex = j
        } else {
            result = strconv.Itoa(cal%10) + result
            extended = int(math.Floor(float64(cal / 10)))
        }
    }

+   if extended != 0 {
+        result = strconv.Itoa(extended) + result
+   }

+   sum = Add(result, sum)
}
```

กรณีที่ input มี 0 นำหน้าจะให้การคำนวนผลลัพธ์ผิด ต้องทำการ handle case ที่ input มี 0 นำ และทำการตัด 0 ข้างหน้าทิ้ง

```diff
for i := lenMultiplier - 1; i >= 0; i-- {
    m1, _ := strconv.Atoi(string(multiplier[i]))

    for z := 0; z < (lenMultiplier - 1 - i); z++ {
        result += "0"
    }

    for j := lenMultiplicand - 1; j >= 0; j-- {
        m2, _ := strconv.Atoi(string(multiplicand[j]))
        cal := m1*m2 + extended

        if m2 == 0 && extended != 0 {
            result = strconv.Itoa(cal) + result
            extended = 0
            lastNonZeroIndex = j
        } else {
            result = strconv.Itoa(cal%10) + result
            extended = int(math.Floor(float64(cal / 10)))
        }
+       if m2 != 0 {
+           lastNonZeroIndex = j
+       }
    }

    if extended != 0 {
        result = strconv.Itoa(extended) + result
    }

+   sum = Add(result[lastNonZeroIndex:], sum)
}
```

