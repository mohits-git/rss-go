package main

import "errors"

type command struct {
	Name string
	Args []string
}

type commands struct {
	registry map[string]func(*state, command) error
}

func (c *commands) register(name string, f func(*state, command) error) {
	c.registry[name] = f
}

func (c *commands) run(s *state, cmd command) error {
	if f, ok := c.registry[cmd.Name]; ok {
		return f(s, cmd)
	}
	return errors.New("Err " + cmd.Name + " command not found")
}

func registerCommands() *commands {
	c := commands{
		registry: make(map[string]func(*state, command) error),
	}

	c.register("login", handleLogin)
	c.register("register", handleRegister)
	c.register("reset", handleReset)
	c.register("users", handleUsers)
	c.register("agg", handleAgg)
	c.register("feeds", handleFeeds)
	c.register("addfeed", middlewareLoggedIn(handleAddFeed))
	c.register("follow", middlewareLoggedIn(handleFollow))
	c.register("following", middlewareLoggedIn(handleFollowing))
	c.register("unfollow", middlewareLoggedIn(handleUnfollow))

	return &c
}
