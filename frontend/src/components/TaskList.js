import React from 'react';
import TaskItem from './TaskItem';
import { Typography, Divider } from '@mui/material';

const TaskList = ({ tasks, onEdit, onDelete }) => {
  return (
    <div>
      {tasks.length === 0 ? (
        <Typography variant="h6" align="center">
          タスクがありません。タスクを追加してください。
        </Typography>
      ) : (
        tasks.map((task) => (
          <div key={task.id}>
            <TaskItem task={task} onEdit={onEdit} onDelete={onDelete} />
            <Divider />
          </div>
        ))
      )}
    </div>
  );
};

export default TaskList;

