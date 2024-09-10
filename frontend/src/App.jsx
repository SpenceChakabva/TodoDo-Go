import React, { useState, useEffect } from 'react';
import TodoList from './components/TodoList';

const App = () => {
  const [todos, setTodos] = useState([]);
  const [newTodo, setNewTodo] = useState('');

  useEffect(() => {
	fetch('/todos')
	  .then(response => response.json())
	  .then(data => setTodos(data));
  }, []);

  const addTodo = () => {
	fetch('/todos', {
	  method: 'POST',
	  headers: {
		'Content-Type': 'application/json',
	  },
	  body: JSON.stringify({ title: newTodo, completed: false }),
	})
	  .then(response => response.json())
	  .then(todo => setTodos([...todos, todo]));
	setNewTodo('');
  };

  const toggleComplete = id => {
	const todo = todos.find(todo => todo.id === id);
	fetch(`/todos/${id}`, {
	  method: 'PUT',
	  headers: {
		'Content-Type': 'application/json',
	  },
	  body: JSON.stringify({ ...todo, completed: !todo.completed }),
	})
	  .then(response => response.json())
	  .then(updatedTodo => {
		setTodos(todos.map(todo => (todo.id === id ? updatedTodo : todo)));
	  });
  };

  const deleteTodo = id => {
	fetch(`/todos/${id}`, {
	  method: 'DELETE',
	}).then(() => {
	  setTodos(todos.filter(todo => todo.id !== id));
	});
  };

  return (
	<div className="container mx-auto p-4">
	  <h1 className="text-2xl font-bold mb-4">Todo List</h1>
	  <div className="flex mb-4">
		<input
		  type="text"
		  className="flex-1 pappearance-none bg-transparent border-none w-full text-gray-700 mr-3 py-1 px-2 leading-tight focus:outline-none-2 border rounded"
		  value={newTodo}
		  onChange={e => setNewTodo(e.target.value)}
		/>
		<button
		  className="bg-blue-500 text-white px-4 py-2 rounded ml-2"
		  onClick={addTodo}
		>
		  Add
		</button>
	  </div>
	  <TodoList todos={todos} toggleComplete={toggleComplete} deleteTodo={deleteTodo} />
	</div>
  );
};

export default App;