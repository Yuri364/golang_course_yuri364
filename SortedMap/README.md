### SortedMap

Сделать пакет `SortedMap` реализующий структуру данных `map`, но с порядком.

Суть этого класса: `map[string]int`, но с сохраненным порядком ключей.
Т.е. мы можем доставать значения из мапы по порядку как их туда записывали.

Использовать его в такой задаче:

Взять [текст] (txt)

Обработать его и найти слова:
- N наиболее употребимых слов из более чем M букв
- не встречающиеся ни в начале ни в конце предложения.
- Вывести в порядке как они встречались в тексте.
- Пунктуацию всю игнорируем (убираем)
- N=10, M=3 для этого текста

[текст]:https://english-e-reader.net/book/gone-with-the-wind-margaret-mitchell

**Результат:**

| #  | Слово  | Кол-во |
|----|--------|--------|
| 1  | like   | 35     |
| 2  | told   | 29     |
| 3  | looked | 39     |
| 4  | marry  | 21     |
| 5  | went   | 62     |
| 6  | love   | 32     |
| 7  | want   | 29     |
| 8  | into   | 29     |
| 8  | took   | 20     |
| 10 | cant   | 18     |