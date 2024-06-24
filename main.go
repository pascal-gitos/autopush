package main

import (
	"bytes"
	"fmt"
	"log"
	"os/exec"
)

func main() {
    // Проверка наличия незакоммиченных изменений
    if hasUncommittedChanges() {
        // Если есть незакоммиченные изменения, коммитим и пушим их
        commitAndPushChanges()
    } else {
        fmt.Println("Нет незакоммиченных изменений")
    }
}

// Проверка наличия незакоммиченных изменений
func hasUncommittedChanges() bool {
    cmd := exec.Command("git", "status", "--porcelain")
    var out bytes.Buffer
    cmd.Stdout = &out
    err := cmd.Run()
    if err != nil {
        log.Fatalf("Ошибка выполнения команды: %v", err)
    }
    // Если вывод не пуст, значит есть незакоммиченные изменения
    return out.Len() > 0
}

// Коммитим и пушим изменения
func commitAndPushChanges() {
    // Команда для добавления всех изменений
    cmdAdd := exec.Command("git", "add", ".")
    err := cmdAdd.Run()
    if err != nil {
        log.Fatalf("Ошибка добавления изменений: %v", err)
    }

    // Команда для коммита изменений
    commitMessage := "Автоматический коммит незакоммиченных изменений"
    cmdCommit := exec.Command("git", "commit", "-m", commitMessage)
    err = cmdCommit.Run()
    if err != nil {
        log.Fatalf("Ошибка коммита изменений: %v", err)
    }

    // Команда для пуша изменений
    cmdPush := exec.Command("git", "push")
    err = cmdPush.Run()
    if err != nil {
        log.Fatalf("Ошибка пуша изменений: %v", err)
    }

    fmt.Println("Изменения успешно закоммичены и запушены")
}