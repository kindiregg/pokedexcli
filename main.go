package main

func main() {
	config := &Config{}
	commands := getCommands(config)
	startRepl(commands)
}
