import axios from 'axios';

const API_BASE_URL = 'http://localhost:8080/api/tasks';

// タスクの取得
export const getTasks = async () => {
  try {
    const response = await axios.get(API_BASE_URL);
    return response.data;
  } catch (error) {
    throw error;
  }
};

// タスクの作成
export const createTask = async (task) => {
  try {
    const response = await axios.post(API_BASE_URL, task);
    return response.data;
  } catch (error) {
    throw error;
  }
};

// タスクの更新
export const updateTask = async (id, updatedTask) => {
  try {
    const response = await axios.put(`${API_BASE_URL}/${id}`, updatedTask);
    return response.data;
  } catch (error) {
    throw error;
  }
};

// タスクの削除
export const deleteTask = async (id) => {
  try {
    const response = await axios.delete(`${API_BASE_URL}/${id}`);
    return response.data;
  } catch (error) {
    throw error;
  }
};

