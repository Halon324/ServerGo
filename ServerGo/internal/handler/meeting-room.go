package handler

import (
	"ServerGo/internal/models"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

func MeetingRoomHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// TODO: создать методы для создания meeting-room, получения по id, получения всех, изменения и удаления
	switch r.Method {
	case http.MethodPost:
		postMeetingRoom(w, r)
	case http.MethodGet:
		getMeetingRoom(w, r)
	case http.MethodDelete:
		deleteMeetingRoom(w, r)
	case http.MethodHead:
		headMeetingRoom(w, r)
	case http.MethodPut:
		putMeetingRoom(w, r)
	}
}

// метод для получения списка MeetingRoom по id
func getMeetingRoom(w http.ResponseWriter, r *http.Request) {
	var mr models.MeetingRoom
	var roomID, err = strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		http.Error(w, "Неверный формат ID:", http.StatusBadRequest)
		return
	}

	// Создаем экземпляр структуры комнаты
	mr = models.MeetingRoom{Id: roomID}

	// Вызываем метод ReadOne() для получения данных конкретной комнаты
	err = mr.ReadOne()
	if err != nil {
		http.Error(w, "Ошибка при чтении данных комнаты:", http.StatusInternalServerError)
		return
	}

	// Формируем ответ
	response, err := json.Marshal(mr)
	if err != nil {
		http.Error(w, "Ошибка при кодировании JSON:", http.StatusInternalServerError)
		return
	}

	// Отправляем ответ клиенту
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

// метод для создания MeetingRoom
func postMeetingRoom(w http.ResponseWriter, r *http.Request) {
	var mr models.MeetingRoom

	// Декодируем JSON-данные из тела запроса
	if err := json.NewDecoder(r.Body).Decode(&mr); err != nil {
		fmt.Println(err.Error())
		http.Error(w, "Ошибка при декодировании JSON:", http.StatusBadRequest)
		return
	}

	// Вызываем метод Create() для создания комнаты
	err := mr.Create()
	if err != nil {
		http.Error(w, "Ошибка при создании комнаты:", http.StatusInternalServerError)
		return
	}

	// Формируем ответ
	response, err := json.Marshal(mr)
	if err != nil {
		http.Error(w, "Ошибка при кодировании JSON:", http.StatusInternalServerError)
		return
	}

	// Отправляем ответ клиенту
	w.WriteHeader(http.StatusCreated)
	w.Write(response)
}

// метод для удаления MeetingRoom
func deleteMeetingRoom(w http.ResponseWriter, r *http.Request) {
	var mr models.MeetingRoom
	// Декодируем JSON-данные из тела запроса
	if err := json.NewDecoder(r.Body).Decode(&mr); err != nil {
		fmt.Println(err.Error())
		http.Error(w, "Ошибка при декодировании JSON:", http.StatusBadRequest)
		return
	}

	// Вызываем метод Delete() для удаления комнаты
	err := mr.Delete()
	if err != nil {
		http.Error(w, "Ошибка при удалении комнаты:", http.StatusInternalServerError)
		return
	}

	// Отправляем ответ клиенту
	w.WriteHeader(http.StatusNoContent) // Успешное удаление без содержимого в ответе
}

// метод для получения всех MeetingRoom
func headMeetingRoom(w http.ResponseWriter, r *http.Request) {
	var mr models.MeetingRoom
	// Декодируем JSON-данные из тела запроса
	if err := json.NewDecoder(r.Body).Decode(&mr); err != nil {
		fmt.Println(err.Error())
		http.Error(w, "Ошибка при декодировании JSON:", http.StatusBadRequest)
		return
	}

	// Вызываем метод Read() для получения данных комнаты
	err := mr.Read()
	if err != nil {
		http.Error(w, "Ошибка при чтении данных комнаты:", http.StatusInternalServerError)
		return
	}
	// Формируем ответ
	response, err := json.Marshal(mr)
	if err != nil {
		http.Error(w, "Ошибка при кодировании JSON:", http.StatusInternalServerError)
		return
	}

	// Отправляем ответ клиенту
	w.WriteHeader(http.StatusOK)
	w.Write(response)

}

// метод для обновления MeetingRoom
func putMeetingRoom(w http.ResponseWriter, r *http.Request) {
	var _ models.MeetingRoom

}

// TODO: сделать API для получения свободных переговорок на данный момент
// TODO: сделать API для бронирования конкретной meeting-room, отмены бронирования
// TODO: ???
// TODO: PROFIT!1!!
