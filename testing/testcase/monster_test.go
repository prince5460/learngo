package testcase

import "testing"

func TestMonster_Store(t *testing.T) {
	monster := Monster{
		Name:  "牛魔王",
		Age:   800,
		Skill: "牛魔拳",
	}

	res := monster.Store()
	if !res {
		t.Fatalf("monster.Store() err,%v", res)
	}
	t.Logf("monster.Store() 测试成功")
}

func TestMonster_ReStore(t *testing.T) {
	var monster Monster
	res := monster.ReStore()
	if !res {
		t.Fatalf("monster.ReStore() err,%v", res)
	}

	if monster.Name != "牛魔王" {
		t.Fatalf("monster.ReStore() err,%v", monster.Name)
	}
	t.Logf("monster.ReStore() 测试成功")
}
