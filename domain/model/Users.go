package domain

import "time"

type Users struct {
    ID        int
    Name      string
    Age       int
    CreatedAt time.Time
}
