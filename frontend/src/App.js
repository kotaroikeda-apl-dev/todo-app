import React, { useState, useEffect } from 'react';
import { Container, Typography, Box } from '@mui/material';
import TaskList from './components/TaskList';
import TaskForm from './components/TaskForm';
import { getTasks, createTask, updateTask, deleteTask } from './api';

function App() {
  const [tasks, setTasks] = useState([]);
  const [taskToEdit, setTaskToEdit] = useState(null); // 編集中のタスク

  // タスクの取得
  const fetchTasks = async () => {
    try {
      const data = await getTasks();
      setTasks(data);
    } catch (error) {
      console.error('タスクの取得中にエラーが発生しました:', error);
    }
  };

  useEffect(() => {
    fetchTasks();
  }, []);

  // 新しいタスクの追加
  const addTask = async (task) => {
    try {
      const newTask = await createTask(task);
      setTasks([...tasks, newTask]);
    } catch (error) {
      console.error('タスクの追加中にエラーが発生しました:', error);
    }
  };

  // タスクの編集を開始
  const editTask = (task) => {
    setTaskToEdit(task);
  };

  // タスクの更新
  const updateExistingTask = async (updatedTask) => {
    try {
      const updated = await updateTask(updatedTask.id, updatedTask);
      setTasks(tasks.map((task) => (task.id === updated.id ? updated : task)));
      setTaskToEdit(null); // 編集終了
    } catch (error) {
      console.error('タスクの更新中にエラーが発生しました:', error);
    }
  };

  // タスクの削除
  const removeTask = async (id) => {
    try {
      await deleteTask(id);
      setTasks(tasks.filter((task) => task.id !== id));
    } catch (error) {
      console.error('タスクの削除中にエラーが発生しました:', error);
    }
  };

  return (
    <Container maxWidth="md">
      <Box sx={{ textAlign: 'center', marginTop: 4, marginBottom: 4 }}>
        <Typography variant="h3" component="h1" gutterBottom>
          タスクマネージャー
        </Typography>
      </Box>
      <TaskForm addTask={addTask} taskToEdit={taskToEdit} updateTask={updateExistingTask} />
      <TaskList tasks={tasks} onEdit={editTask} onDelete={removeTask} />
    </Container>
  );
}

export default App;

