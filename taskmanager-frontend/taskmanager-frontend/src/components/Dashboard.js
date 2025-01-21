import React, { useEffect, useState } from "react";
import { fetchTasks, createTask, deleteTask } from "../api";
import {
  Container,
  Typography,
  Button,
  TextField,
  List,
  ListItem,
  IconButton,
} from "@mui/material";
import DeleteIcon from "@mui/icons-material/Delete";

const Dashboard = () => {
  const [tasks, setTasks] = useState([]);
  const [newTask, setNewTask] = useState("");

  const loadTasks = async () => {
    const { data } = await fetchTasks();
    setTasks(data);
  };

  const handleAddTask = async () => {
    if (!newTask.trim()) return;
    await createTask({ title: newTask });
    setNewTask("");
    loadTasks();
  };

  const handleDeleteTask = async (id) => {
    await deleteTask(id);
    loadTasks();
  };

  useEffect(() => {
    loadTasks();
  }, []);

  return (
    <Container>
      <Typography variant="h4" gutterBottom>
        Task Dashboard
      </Typography>
      <TextField
        fullWidth
        label="New Task"
        value={newTask}
        onChange={(e) => setNewTask(e.target.value)}
      />
      <Button variant="contained" color="primary" onClick={handleAddTask}>
        Add Task
      </Button>
      <List>
        {tasks.map((task) => (
          <ListItem key={task.id}>
            {task.title}
            <IconButton edge="end" onClick={() => handleDeleteTask(task.id)}>
              <DeleteIcon />
            </IconButton>
          </ListItem>
        ))}
      </List>
    </Container>
  );
};

export default Dashboard;
