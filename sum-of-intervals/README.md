**Codewars Problem: Sum of Intervals**

https://www.codewars.com/kata/52b7ed099cdc285c300001cd

**Codewars Result**

![result](https://github.com/D4rk1ink/Codewars-Problems/blob/master/sum-of-intervals/result.png?raw=true)

**Problem**

ให้หาผลรวมของ intervals ใน array[][2] ทั้งหมด โดย interval ที่ overlap กันจะนับเป็น 1 interval

```javascript
// Input
[[1,5]]

// Output
// 4
```
```javascript
// Input
[[1,5],[7,10]]

// Output
// (4 + 3) = 7
```
```javascript
// Input
[[1,5],[7,10],[9,12]]

// Output
// after merged = [[1,5],[7,12]] = 9
```

**Solution**

เนื่องด้วย intervals array ไม่มีการ sort มาให้ จึงต้อง sort intervals โดยเรียงละดับจาก array[][1] เพื่อที่จะได้หา interval ที่มีการ overlap กันง่ายขึ้น
```go
// SortIntervals Function

for i, interval := range intervals {
    if intervals[pivot][0] <= interval[0] {
        temp := append([][2]int{}, intervals[0:i]...)
        temp = append(temp, intervals[pivot])
        temp = append(temp, intervals[i:pivot]...)
        temp = append(temp, intervals[pivot+1:]...)
        intervals = temp
        break
    }
}
```
จาก code จะทำการจับ interval ทุกตัวมาวนหาลำดับใหม่ โดยจะหาว่า interval ที่ focus ควรอยู่หน้า interval ตัวไหนและจะทำการลำดับ array ใหม่

เมื่อได้ sorting ของ intervals มาแล้วก็จะเริ่มวนนับช่วงในแต่ละ interval
```go
start := sortedIntervals[0][0]
end := sortedIntervals[0][1]

for _, interval := range sortedIntervals {
    if end < interval[0] {
        sum = sum + (end - start)
        start = interval[0]
        end = interval[1]
    } else if interval[1] > end {
        end = interval[1]
    }
}
```
จาก code จะเริ่มจากกำหนด start end pivot ไว้เพื่อจะ slide pivot นับช่วงไปทุก ๆ intervals

เริ่มจากกำหนด start end จาก interval อันแรกสุด กรณีที่วนเจอ interval ที่ overlap กับ start และ end จะทำการ update end ให้มีค่าที่มีช่วงมากขึ้น

กรณีเจอ interval ที่ไม่ได้ overlap กับ start และ end จะทำการนับช่วงของ start และ end เก็บไว้และเริ่มกำหนดค่า start และ end เป็น interval ล่าสุดใหม่ ทำแบบนี้ไปเรื่อย ๆ จนครบก็จะได้ผลรวมของ intervals ทั้งหมด

**Improvement**

cost ที่ใช้ในการ sort ยังรู้สึกค่อนข้างมาก ถ้าเลี่ยงการใช้ sort จะช่วยลด cost ได้ process ลดได้

statement ที่ทำการ order array เป็นการสร้าง array ขึ้นมาใหม่ ทำให้มีการ cost เรื่อง memory เพิ่มขึ้น

**Searching Reference**

[go-slice-trick](https://ueokande.github.io/go-slice-tricks/)