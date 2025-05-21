**Problem**

create function ในการ reverse array โดย code size ต้องไม่เกิน 30 characters

เริ่มต้น code จาก
```javascript
reverse=a=>a.map((_,i)=>a[a.length-1-i])
```
และ

```javascript
reverse=a=>a.map(_=>a.pop())
```
จาก code ก่อน enhance code size ได้ แต่ output ไม่ถูกต้อง เพราะมีการ loop และ pop ออกไปเรื่อย ๆ พร้อม ๆ กัน ทำให้ loop รอบหลัง ๆ จะได้ empty เพราะการ pop array เป็นการ mutate ระดับ address (reference by address)

**AI Guideline**

Final Solution from AI guideline
```javascript
reverse=a=>[...a].map(a.pop,a)
```

ทำการแยก array ที่จะทำการ loop และ pop ออกจากกัน โดย `[...a]` จะใช้ในการ loop และใช้ array ที่รับมาในการ pop แต่ใน code จะใช้ function pop ในการเป็น callback ของ function map ซึ่ง by default พวก array function จะมีการ binding `this` กับ array นั้น ๆ แต่ใน code จะเป็นการ binding `this` ของ array ที่ใช้ในการ pop เพื่อไม่ให้ array ของ loop และ pop ชนกัน

