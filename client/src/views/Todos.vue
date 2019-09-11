<template>
<div id="todo-list">
  <h1>閉鎖されました</h1>
  <!--h1 id="title">{{ user }}'s Todos</h1>
  <div>
    <input type="text" name="name" v-model="newTodo" placeholder="new todo">
    <button type="button" name="add" @click="postTodo">Add</button>
  </div>
  <div v-if="todos.length === 0">
    <p>No your todos...</p>
  </div>
  <ul v-else id="todo-list">
    <li v-for="(todo, index) in todos" :key="index">
      <input type="checkbox" :id="index" name="completed" :checked="todo.completed" @click="checkTodo(todo.id)">
      <label :for="index">{{ todo.name }}</label>
      <button type="button" name="delete" @click="deleteTodo(todo.id)">X</button>
    </li>
  </ul>
  <button type="button" name="logout" @click="logout()">logout</button-->
</div>
</template>

<script lang="ts">
/*
import { Component, Vue } from 'vue-property-decorator';

@Component({
  components: {
  }
})
export default class Todos extends Vue {
  
  private user: string = '';
  private newTodo: string = '';
  private todos: string[] = [];
  private completedTodos: string[] = []; 
  
  private created(): void {
    const date = new Date();
    const claims = JSON.parse(atob(this.getToken().split('.')[1]));
    this.user = claims.name;
    if(claims.exp < Math.floor(date.getTime() / 1000)) {
      this.logout();
    } else {
      this.getTodos();
    }
  }
  
  private getTodos(): void {
    const url = 'api/todos';
    const headers = {'Authorization': `Bearer ${this.getToken()}`};

    fetch(url, {headers}).then(response => {
      if(response.ok) {
        return response.json();
      }
      return [];
    }).then(json => {
      this.todos = json;
    })
  }
  private postTodo(): void {
    const url = 'api/todos';
    const method = 'POST';
    const headers = {
      'Authorization': `Bearer ${this.getToken()}`,;
      'Content-Type': 'application/json; charset=UTF-8';
    }
    const body = JSON.stringify({name: this.newTodo});

    fetch(url, {method, headers, body}).then(response => {
      if(response.ok) {
        return response.json();
      }
    }).then(json => {
      if(typeof json === 'undefined') {
        return;
      }
      this.todos.push(json);
      this.newTodo = '';
    });
  }
  private deleteTodo(id: number): void {
    const url = `api/todos/${id}`;
    const method = 'DELETE';
    const headers = {'Authorization': `Bearer ${this.getToken()}`};

    fetch(url, {method, headers}).then(response => {
      if(response.ok) {
        this.todos = this.todos.filter(todo => todo.id !== id);
      }
    })
  }
  private checkTodo(id: number): void {
    const url = `api/todos/${id}/completed`;
    const method = 'PUT';
    const headers = {'Authorization': `Bearer ${this.getToken()}`};

    fetch(url, {method, headers});
  }
  private getToken: any {
    return localStorage.getItem('token');
  }
  private logout(): void {
    localStorage.removeItem('token');
  }

}*/
</script>
