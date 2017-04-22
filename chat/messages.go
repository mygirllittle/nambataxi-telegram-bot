package chat

const (
	BOT_WELCOME_MESSAGE =
`Вас приветствует Telegram-бот для вызова NambaTaxi!

Заказывайте такси с моей помощью за считанные секунды:
1. Вызовите команду "Быстрый заказ такси" кнопкой снизу, либо вводом текста
2. Введите телефон
3. Выберите тариф
4. Введите адрес

Я умный (и продолжаю учиться), поэтому после первой поездки запоминаю введенные номера и адреса, и при повторном заказе вам не придется вводить их заново. В процессе обработки заказа, вы можете узнать его статус.

Приятного использования!`

	BOT_ORDER_DONE =
`Ваш заказ выполнен! Спасибо, что воспользовались услугами Намба Такси. Если вдруг что-то не так, то телефон Отдела Контроля Качества к вашим услугам:
+996 (312) 97-90-60
+996 (701) 97-67-03
+996 (550) 97-60-23`

	BOT_FARES_LINK = "Для получения более подробной информации, перейдите по ссылке: https://nambataxi.kg/ru/tariffs/"
	BOT_NO_ORDERS = "К сожалению у вас нет заказа"
	BOT_ERROR_GET_FARES = "Ошибка. Не удалось получить тарифы. Попробуйте еще раз"
	BOT_ASK_PHONE = "Укажите ваш телефон. Например: +996555112233\n\n/Cancel"
	BOT_PHONE_START_996 = "Телефон должен начинаться с +996"
	BOT_ASK_FARE = "Телефон сохранен. Теперь укажите тариф\n\n/Cancel"
	BOT_ASK_ADDRESS = "Укажите ваш адрес. Куда подать машину?\n\n/Cancel"
	BOT_ERROR_GET_1_FARE = "Ошибка! Не удалось получить тариф по имени. Попробуйте еще раз"
	BOT_ERROR_EARLY_NEAREST_DRIVERS = "Для начала нужно создать заказ"

)
