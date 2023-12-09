package runner

import (
	"fmt"
	"os"
	"os/exec"
)

func checkApp(path string) (err error) {
	// Проверка существования файла
	_, err = os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			return fmt.Errorf("файл %s не существует", path)
		}
		return fmt.Errorf("ошибка при проверке файла: %v", err)
	}

	/*// Проверка является ли файл исполняемым
	if runtime.GOOS == "windows" {
		return checkWindowsApp(path)
	} else {
		return checkUnixApp(path)
	}*/
	return err
}

func checkWindowsApp(path string) error {
	cmd := exec.Command("where", path)
	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("файл %s не является исполняемым: %v", path, err)
	}
	return nil
}

func checkUnixApp(path string) error {
	cmd := exec.Command("sh", "-c", fmt.Sprintf("stat -c '%%a' %s | grep 'x'", path))
	err := cmd.Run()
	if err != nil {
		fmt.Println(err)
		fmt.Println("+++")
		return fmt.Errorf("файл %s не является исполняемым: %v", path, err)
	}
	return nil
}

func RunApp(pathApp string) (err error) {

	err = checkApp(pathApp)
	if err != nil {
		return err
	}

	cmd := exec.Command(pathApp)
	/*	defer func() {
		err := cmd.Process.Kill()
		if err != nil {
			fmt.Println("Error killing process:", err)
		}
		fmt.Println("1.exe killed")
		fmt.Printf(" app %v killed \n", pathApp)
	}()*/
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err = cmd.Start()
	if err != nil {
		fmt.Println("Error starting executable:", err)
		return err
	}
	fmt.Printf("Run app %v \n", pathApp)

	return err
}
