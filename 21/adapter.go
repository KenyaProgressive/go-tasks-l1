package main

import (
	"fmt"
	"strconv"
	"time"
)

type MessageSendAdapter interface { // Клиент использует данныый метод для отправки сообщения
	SendMessage(id int)
}

// Telegram ===============================================

type TelegramUser struct{} // Телеграм-клиент

func (tUser *TelegramUser) sendTelegramMessage(toSentId int) { // Метод для отправки сообщения в Телеграме
	if len(strconv.Itoa(toSentId)) == 8 { // валидация id
		time.Sleep(time.Millisecond * 20) // уникальное время отправки
		fmt.Println("Telegram message was sent")
	} else {
		fmt.Println("Incorrect Telegram Id")
	}
}

type TelegramUserAdapter struct { // Создаем адаптер для телеграм-клиента -- сохраняем ссылку на оригинал объекта
	*TelegramUser
}

func (TelegramAdapter *TelegramUserAdapter) SendMessage(id int) { // реализуем метод клиентского интерфейса в адаптере
	TelegramAdapter.sendTelegramMessage(id)
}

func NewTelegramAdapter(tUser *TelegramUser) MessageSendAdapter { // конструктор адаптера для Телеграм-Клиента
	return &TelegramUserAdapter{tUser}
}

// ===============================================================

// WhatsApp ======================================================

type WhatsAppUser struct{} // WhatsApp-клиент

func (whUser *WhatsAppUser) sendWhatsAppMessage(toSentId int) {
	if len(strconv.Itoa(toSentId)) == 6 { // валидация id
		time.Sleep(time.Millisecond * 50) // уникальное время отправки
		fmt.Println("WhatsApp message was sent")
	} else {
		fmt.Println("Incorrect WhatsApp Id")
	}
}

type WhatsAppUserAdapter struct { // Создаем адаптер для ватсап-клиента -- сохраняем ссылку на оригинал объекта
	*WhatsAppUser
}

func (whAdapter *WhatsAppUserAdapter) SendMessage(id int) { // реализуем метод клиентского интерфейса в адаптере
	whAdapter.sendWhatsAppMessage(id)
}

func NewWhatsAppAdapter(whUser *WhatsAppUser) MessageSendAdapter { // конструктор адаптера для ватсап-Клиента
	return &WhatsAppUserAdapter{whUser}
}

// ===============================================================

// VK ============================================================

type VKUser struct{} // VK-клиент

func (vkUser *VKUser) sendVKMessage(toSentId int) {
	if len(strconv.Itoa(toSentId)) == 11 { // валидация id
		time.Sleep(time.Millisecond * 30) // уникальное время отправки
		fmt.Println("VK message was sent")
	} else {
		fmt.Println("Incorrect VK Id")
	}
}

type VKUserAdapter struct { // Создаем адаптер для VK-клиента -- сохраняем ссылку на оригинал объекта
	*VKUser
}

func (vkAdapter VKUserAdapter) SendMessage(id int) { // реализуем метод клиентского интерфейса в адаптере
	vkAdapter.sendVKMessage(id)
}

func NewVkAdapter(vkUser *VKUser) MessageSendAdapter { // конструктор адаптера для VK-Клиента
	return &VKUserAdapter{vkUser}
}

// ===============================================================
