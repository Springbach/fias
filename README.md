
# address


## Indices

  * [houses](#1-houses)
  * [territory](#2-territory)
  * [streets](#3-streets)
  * [regions](#4-regions)
  * [cities](#5-cities)

--------


### 1. houses


retrieve house array


***Endpoint:***

```bash
Method: GET
Type: RAW
URL: {{url}}/list/houses
```


***Query params:***

| Key | Value | Description |
| --- | ------|-------------|
| zip | {{zip}} |  |
| street | {{street}} |  |
|  |  |  |


***Responses:***


Status: houses | Code: 200


```js
{
    "hs": [
        {
            "bn": "",
            "cn": "",
            "div": "0",
            "hn": "5",
            "sn": ""
        },
        {
            "bn": "",
            "cn": "",
            "div": "0",
            "hn": "27",
            "sn": ""
        },
        {
            "bn": "",
            "cn": "",
            "div": "0",
            "hn": "16",
            "sn": ""
        },
        {
            "bn": "",
            "cn": "",
            "div": "0",
            "hn": "21",
            "sn": ""
        },
        {
            "bn": "",
            "cn": "",
            "div": "0",
            "hn": "2",
            "sn": ""
        },
        {
            "bn": "",
            "cn": "",
            "div": "0",
            "hn": "16/3",
            "sn": ""
        }
    ]
}
```


### 2. territory


retrieve streets array


***Endpoint:***

```bash
Method: GET
Type: RAW
URL: {{url}}/list/territory
```


***Query params:***

| Key | Value | Description |
| --- | ------|-------------|
| zip | {{zip}} |  |


***Responses:***


Status: territory | Code: 200


```js
{
    "city": "г.Новосибирск",
    "streets": [
        "проезд.Пожарского",
        "рзд.Иня 12 км",
        "ул.Артема",
        "ул.Белая",
        "ул.Героев Революции",
        "ул.Заречная",
        "ул.Зеркальная",
        "ул.Красный Факел",
        "ул.Марата",
        "ул.Молодогвардейская",
        "ул.Первомайская",
        "ул.Перова",
        "ул.Подольская",
        "ул.Пожарского",
        "ул.Поленова",
        "ул.Прохладная",
        "ул.Радиаторная",
        "ул.Серебристая",
        "ул.Столбовая",
        "ул.Чапаева",
        "ул.Электровозная",
        "ш.Бердское"
    ]
}
```


### 3. streets


retrieve streets array


***Endpoint:***

```bash
Method: GET
Type: RAW
URL: {{url}}/list/streets
```


***Query params:***

| Key | Value | Description |
| --- | ------|-------------|
| reg | {{reg}} |  |
| city | {{city}} |  |
|  |  |  |


***Responses:***


Status: streets | Code: 200


```js
{
    "streets": [
        "ул.Загородная",
        "ул.Закавказская",
        "ул.Закарпатская",
        "ул.Залесского",
        "ул.Заобская",
        "ул.Заозерная",
        "ул.Западная",
        "ул.Зареченская",
        "ул.Заречная",
        "ул.Заслонова",
        "ул.Затонная",
        "ул.Звездная",
        "ул.Звенигородская",
        "ул.Здвинская",
        "ул.Зейская",
        "ул.Зеленая",
        "ул.Зеленая (п Озерный)",
        "ул.Зеленая Горка",
        "ул.Зеленодолинская",
        "ул.Зеленхозовская",
        "ул.Земнухова",
        "ул.Зенитная",
        "ул.Зеркальная",
        "ул.Зимняя",
        "ул.Златоустовская",
        "ул.Знаменская",
        "ул.Зои Космодемьянской",
        "ул.Золотодолинская",
        "ул.Золоторожская",
        "ул.Зональная",
        "ул.Зоологическая",
        "ул.Зорге",
        "ул.Зыряновская",
        "ул.Иванова",
        "ул.Ивлева",
        "ул.Игарская",
        ...
  
    
    ]
}
```


### 4. regions


retrieve regions array


***Endpoint:***

```bash
Method: GET
Type: RAW
URL: {{url}}/list/regions
```


***Responses:***


Status: regions | Code: 200


```js
{
    "regions": [
        "01",
        "02",
        "03",
        "04",
        "05",
        "06",
        ...
    ]
}
```


### 5. cities


retrieve cities array


***Endpoint:***

```bash
Method: GET
Type: RAW
URL: {{url}}/list/cities
```


***Query params:***

| Key | Value | Description |
| --- | ------|-------------|
| reg | {{reg}} |  |


***Responses:***


Status: cities | Code: 200


```js
{
    "cities": [
        "г.Барабинск",
        "г.Бердск",
        "г.Болотное",
        "г.Искитим",
        "г.Карасук",
        "г.Каргат",
        "г.Куйбышев",
        "г.Купино",
        "г.Новосибирск",
        "г.Обь",
        "г.Татарск",
        "г.Тогучин",
        "г.Черепаново",
        "г.Чулым",
        "г.Чулым-3"
    ]
}
```

---
[Back to top](#address)
