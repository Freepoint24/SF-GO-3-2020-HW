# Код из скринкаста к заданию 11.5.1

[Оригинальный код](https://lms.skillfactory.ru/asset-v1:SkillFactory+GO+30SEPT2020+type@asset+block@GO_11.5_ring.go)

## Паника
Метод `Read` вызывает панику (обращение по индексу за границами массива)
в случае если последовательно произвести операции записи и чтения из буфера, ё
больше чем размер буфера

#### Воспроизведение
`go test .`

Результат:
```
--- FAIL: TestIntRing_Read (0.00s)
panic: runtime error: index out of range [5] with length 5 [recovered]
        panic: runtime error: index out of range [5] with length 5

goroutine 6 [running]:
testing.tRunner.func1.2(0x1135360, 0xc000014090)
        testing.go:1143 +0x332
testing.tRunner.func1(0xc000001380)
        testing.go:1146 +0x4b6
panic(0x1135360, 0xc000014090)
        panic.go:965 +0x1b9
sf-golang-11-5-1.(*IntRing).Read(0xc000010740, 0x2a, 0x0, 0x0)
        main.go:66 +0x3ad
sf-golang-11-5-1.TestIntRing_Read(0xc000001380)
        main_test.go:19 +0x139
testing.tRunner(0xc000001380, 0x114e630)
        testing.go:1193 +0xef
created by testing.(*T).Run
        testing.go:1238 +0x2b3
FAIL    sf-golang-11-5-1        0.201s
FAIL

```

## Примечание
Методы `Read` и `Print` реализованы крайне неоптимально с тз вычислительной сложности, тк для работы используют методы `IsEmpty` и `getContinuousArray`, которые в самом худшем случае делают фулл-скан всей коллекции — т.е. кратно увеличивают сложность.
