package models

import (
	"ServerGo/internal/sqlite"
	"database/sql"
	"errors"
	"fmt"
	"time"
)

type Device struct {
	Name string `json:"name"`
	Type string `json:"type"`
}

type MeetingRoom struct {
	Id           int64         `json:"id"`
	Number       string        `json:"number"`
	Name         string        `json:"name"`
	Places       int32         `json:"places"`
	Device       int64         `json:"device"`
	WorkingHours time.Duration `json:"working_hours"`
}

type MeetingRoomSchedule struct {
	MeetingRoom MeetingRoom   `json:"meeting_room"`
	Duration    time.Duration `json:"duration"`
	Persons     []Person      `json:"persons"`
}

func (mr *MeetingRoom) Read() error {
	result, err := sqlite.DB.Query("SELECT * FROM meeting_room")
	if err != nil {
		return fmt.Errorf("ошибка при выполнении запроса: %w", err)
	}
	defer func(result *sql.Rows) {
		err := result.Close()
		if err != nil {

		}
	}(result)

	// Чтение данных из таблицы
	for result.Next() {
		var meetingRoom MeetingRoom
		err := result.Scan(
			&meetingRoom.Id, // Если поле ID есть в таблице
			&meetingRoom.Number,
			&meetingRoom.Name,
			&meetingRoom.Places,
			&meetingRoom.Device,
			&meetingRoom.WorkingHours,
		)
		if err != nil {
			return fmt.Errorf("ошибка при чтении данных: %w", err)
		}

		fmt.Println(result.ColumnTypes())
	}

	// Обработка ошибок при чтении данных
	if err := result.Err(); err != nil {
		return fmt.Errorf("ошибка при чтении данных: %w", err)
	}
	return nil
}
func (mr *MeetingRoom) Create() error {
	result, err := sqlite.DB.Prepare("INSERT INTO MeetingRoom (number, name, places, devices, workinghours) VALUES ($1, $2, $3,$4,$5)")
	if err != nil {
		return fmt.Errorf("ошибка при подготовке запроса: %w", err)
	}
	defer func(result *sql.Stmt) {
		err := result.Close()
		if err != nil {

		}
	}(result)

	// Выполнение запроса
	stmt, err := result.Exec(mr.Number, mr.Name, mr.Places, mr.Device, mr.WorkingHours)
	if err != nil {
		return fmt.Errorf("ошибка при выполнении запроса: %w", err)
	}

	// Получение ID созданной записи
	lastInsertId, err := stmt.LastInsertId()
	if err != nil {
		return fmt.Errorf("ошибка при получении ID: %w", err)
	}
	mr.Id = int64(int(lastInsertId)) // Обновление ID в структуре
	return nil
}
func (mr *MeetingRoom) Delete() error {
	result, err := sqlite.DB.Exec("DELETE FROM meeting_room WHERE id = $1", mr.Id)
	if err != nil {
		return fmt.Errorf("ошибка при подготовке запроса: %w", err)
	}
	fmt.Println(result.LastInsertId())
	return nil
}
func (mr *MeetingRoom) Update() error {
	result, err := sqlite.DB.Exec("UPDATE MeetingRoom (number, name, places, devices, workinghours) VALUES ($1, $2, $3,$4,$5), WHERE id = $6", mr.Number, mr.Name, mr.Places, mr.Device, mr.WorkingHours, mr.Id)
	if err != nil {
		return fmt.Errorf("ошибка при подготовке запроса: %w", err)
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("ошибка при получении количества измененных строк: %w", err)
	}
	if rowsAffected == 0 {
		return fmt.Errorf("запись с id %d не найдена", mr.Id)
	}
	fmt.Println(result.LastInsertId())
	return nil
}
func (mr *MeetingRoom) ReadOne() error {
	result, err := sqlite.DB.Query("SELECT * FROM meeting_room WHERE id=$1", mr.Id)
	if err != nil {
		return fmt.Errorf("ошибка при подготовке запроса: %w", err)
	}
	defer func(result *sql.Rows) {
		err := result.Close()
		if err != nil {

		}
	}(result)

	// Сканирование результата
	err = result.Scan(
		&mr.Number,
		&mr.Name,
		&mr.Places,
		&mr.Device,
		&mr.WorkingHours,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return fmt.Errorf("запись с id %d не найдена", mr.Id)
		}
		return fmt.Errorf("ошибка при чтении данных: %w", err)
	}
	fmt.Println(result.ColumnTypes())
	return nil
}

//var meetingRooms := MeetingRoom{}
