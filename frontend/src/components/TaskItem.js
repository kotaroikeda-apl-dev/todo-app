import React from 'react';
import { Card, CardContent, Typography, Button, Box } from '@mui/material';
import EditIcon from '@mui/icons-material/Edit';
import DeleteIcon from '@mui/icons-material/Delete';

const TaskItem = ({ task, onEdit, onDelete }) => {
  return (
    <Card sx={{ marginBottom: 2 }}>
      <CardContent>
        <Typography variant="h6" component="div">
          {task.title}
        </Typography>
        <Typography variant="body2" color="text.secondary" sx={{ marginBottom: 1 }}>
          {task.description}
        </Typography>
        <Typography variant="body2" color="text.secondary" sx={{ marginBottom: 2 }}>
          ステータス: {task.completed ? '完了' : '未完了'}
        </Typography>
        <Box sx={{ display: 'flex', gap: 1 }}>
          <Button
            variant="outlined"
            color="primary"
            startIcon={<EditIcon />}
            onClick={() => onEdit(task)}
          >
            編集
          </Button>
          <Button
            variant="outlined"
            color="error"
            startIcon={<DeleteIcon />}
            onClick={() => onDelete(task.id)}
          >
            削除
          </Button>
        </Box>
      </CardContent>
    </Card>
  );
};

export default TaskItem;

