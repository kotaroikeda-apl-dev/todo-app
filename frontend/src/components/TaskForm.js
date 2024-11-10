import React, { useState, useEffect } from 'react';
import { TextField, Button, Box, Paper, Typography } from '@mui/material';

const TaskForm = ({ addTask, taskToEdit, updateTask }) => {
  const [title, setTitle] = useState('');
  const [description, setDescription] = useState('');

  useEffect(() => {
    if (taskToEdit) {
      setTitle(taskToEdit.title);
      setDescription(taskToEdit.description);
    } else {
      setTitle('');
      setDescription('');
    }
  }, [taskToEdit]);

  const handleSubmit = (e) => {
    e.preventDefault();
    if (!title.trim() || !description.trim()) {
      alert('タイトルと説明を入力してください。');
      return;
    }

    if (taskToEdit) {
      // 編集モード
      updateTask({ ...taskToEdit, title, description });
    } else {
      // 追加モード
      addTask({ title, description, completed: false });
    }

    // フォームをリセット
    setTitle('');
    setDescription('');
  };

  return (
    <Paper elevation={3} sx={{ padding: 3, marginBottom: 4 }}>
      <Typography variant="h5" component="h2" gutterBottom>
        {taskToEdit ? 'タスクを編集' : '新しいタスクを追加'}
      </Typography>
      <Box component="form" onSubmit={handleSubmit}>
        <TextField
          label="タスクのタイトル"
          variant="outlined"
          fullWidth
          value={title}
          onChange={(e) => setTitle(e.target.value)}
          required
          sx={{ marginBottom: 2 }}
        />
        <TextField
          label="タスクの説明"
          variant="outlined"
          fullWidth
          multiline
          rows={4}
          value={description}
          onChange={(e) => setDescription(e.target.value)}
          required
          sx={{ marginBottom: 2 }}
        />
        <Button variant="contained" color="primary" type="submit" fullWidth>
          {taskToEdit ? 'タスクを更新' : 'タスクを追加'}
        </Button>
      </Box>
    </Paper>
  );
};

export default TaskForm;

