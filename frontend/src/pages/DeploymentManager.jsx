import { useEffect, useState } from "react";

import {
  Box,
  Paper,
  Typography,
  Button,
  Table,
  TableHead,
  TableRow,
  TableCell,
  TableBody,
  Dialog,
  DialogTitle,
  DialogContent,
  DialogActions,
  TextField,
  Snackbar,
  Alert,
  Chip,
  CircularProgress,
} from "@mui/material";

import {
  Delete,
  RestartAlt,
  Add,
  OpenInFull,
  Refresh,
} from "@mui/icons-material";

import {
  getDeployments,
  createDeployment,
  deleteDeployment,
  restartDeployment,
  scaleDeployment,
} from "../services/api";

export default function DeploymentManager() {

    const [deployments, setDeployments] = useState([]);
    const [loading, setLoading] = useState(false);

    const [deploymentName, setDeploymentName] = useState("");
    const [image, setImage] = useState("");
    const [replicas, setReplicas] = useState(1);
    const [port, setPort] = useState(80);

    const [openCreate, setOpenCreate] = useState(false);

    const [snackbar, setSnackbar] = useState({
    open: false,
    message: "",
    severity: "success",
    });

    const loadDeployments = async () => {
    try {
        setLoading(true);

        const res = await getDeployments();

        setDeployments(res.data);
    } catch (err) {
        console.error(err);

        showMessage("Failed to load deployments", "error");
    } finally {
        setLoading(false);
    }
    };
    const handleDelete = async (name) => {
    if (!window.confirm(`Delete deployment "${name}"?`)) {
        return;
    }

    try {
        await deleteDeployment(name);

        showMessage("Deployment deleted");

        loadDeployments();
    } catch (error) {
        console.error(error);

        showMessage("Delete failed", "error");
    }
    };
    const handleRestart = async (name) => {
    try {
        await restartDeployment(name);

        showMessage("Deployment restarted");

        loadDeployments();
    } catch (error) {
        console.error(error);

        showMessage("Restart failed", "error");
    }
    };
    const handleScale = async (name) => {
    const value = prompt("Enter replicas");

    if (value === null) return;

    const replicas = parseInt(value);

    if (isNaN(replicas) || replicas < 1) {
        showMessage("Invalid replica count", "error");
        return;
    }

    try {
        await scaleDeployment(name, replicas);

        showMessage("Deployment scaled");

        loadDeployments();
    } catch (error) {
        console.error(error);

        showMessage("Scaling failed", "error");
    }
    };
    const handleCreate = async () => {
    try {
        await createDeployment({
        namespace: "default",
        name: deploymentName,
        image: image,
        replicas: replicas,
        port: port,
        });

        showMessage("Deployment created successfully");

        setDeploymentName("");
        setImage("");
        setReplicas(1);
        setPort(80);

        setOpenCreate(false);

        loadDeployments();
    } catch (error) {
        console.error(error);

        showMessage("Failed to create deployment", "error");
    }
    };
    const showMessage = (message, severity = "success") => {
    setSnackbar({
        open: true,
        message,
        severity,
    });
    };

    useEffect(() => {

        loadDeployments();

    }, []);

    return (
    <>
        <Box p={3}>
        <Paper sx={{ p: 3 }}>

            <Box
            display="flex"
            justifyContent="space-between"
            alignItems="center"
            mb={3}
            >
            <Typography variant="h4">
                Deployment Manager
            </Typography>

            <Box>

                <Button
                variant="outlined"
                startIcon={<Refresh />}
                onClick={loadDeployments}
                sx={{ mr: 2 }}
                >
                Refresh
                </Button>

                <Button
                variant="contained"
                startIcon={<Add />}
                onClick={() => setOpenCreate(true)}
                >
                New Deployment
                </Button>

            </Box>
            </Box>

            {loading ? (
            <Box textAlign="center" py={5}>
                <CircularProgress />
            </Box>
            ) : (
            <Table>

                <TableHead>

                <TableRow>

                    <TableCell>Name</TableCell>

                    <TableCell>Image</TableCell>

                    <TableCell>Replicas</TableCell>

                    <TableCell>Status</TableCell>

                    <TableCell>Available</TableCell>

                    <TableCell align="center">
                    Actions
                    </TableCell>

                </TableRow>

                </TableHead>

                <TableBody>

                {deployments.map((deployment) => (

                    <TableRow key={deployment.name}>

                    <TableCell>
                        {deployment.name}
                    </TableCell>

                    <TableCell>
                        {deployment.image}
                    </TableCell>

                    <TableCell>
                        {deployment.replicas}
                    </TableCell>

                    <TableCell>

                        {deployment.readyReplicas === deployment.replicas ? (

                        <Chip
                            label="Running"
                            color="success"
                        />

                        ) : (

                        <Chip
                            label="Pending"
                            color="warning"
                        />

                        )}

                    </TableCell>

                    <TableCell>
                        {deployment.availableReplicas}
                    </TableCell>

                    <TableCell>

                        <Button
                        size="small"
                        variant="contained"
                        startIcon={<OpenInFull />}
                        sx={{ mr: 1, mb: 1 }}
                        onClick={() => handleScale(deployment.name)}
                        >
                        Scale
                        </Button>

                        <Button
                        size="small"
                        color="success"
                        variant="contained"
                        startIcon={<RestartAlt />}
                        sx={{ mr: 1, mb: 1 }}
                        onClick={() => handleRestart(deployment.name)}
                        >
                        Restart
                        </Button>

                        <Button
                        size="small"
                        color="error"
                        variant="contained"
                        startIcon={<Delete />}
                        sx={{ mb: 1 }}
                        onClick={() => handleDelete(deployment.name)}
                        >
                        Delete
                        </Button>

                    </TableCell>

                    </TableRow>

                ))}

                </TableBody>

            </Table>
            )}

        </Paper>

        </Box>

        <Dialog
        open={openCreate}
        onClose={() => setOpenCreate(false)}
        maxWidth="sm"
        fullWidth
        >

        <DialogTitle>
            Create Deployment
        </DialogTitle>

        <DialogContent>

            <TextField
            label="Deployment Name"
            fullWidth
            margin="normal"
            value={deploymentName}
            onChange={(e) => setDeploymentName(e.target.value)}
            />

            <TextField
            label="Docker Image"
            fullWidth
            margin="normal"
            value={image}
            onChange={(e) => setImage(e.target.value)}
            />

            <TextField
            label="Replicas"
            type="number"
            fullWidth
            margin="normal"
            value={replicas}
            onChange={(e) => setReplicas(Number(e.target.value))}
            />

            <TextField
            label="Port"
            type="number"
            fullWidth
            margin="normal"
            value={port}
            onChange={(e) => setPort(Number(e.target.value))}
            />

        </DialogContent>

        <DialogActions>

            <Button
            onClick={() => setOpenCreate(false)}
            >
            Cancel
            </Button>

            <Button
            variant="contained"
            onClick={handleCreate}
            >
            Create
            </Button>

        </DialogActions>

        </Dialog>

        <Snackbar
        open={snackbar.open}
        autoHideDuration={3000}
        onClose={() =>
            setSnackbar({
            ...snackbar,
            open: false,
            })
        }
        >
        <Alert severity={snackbar.severity}>
            {snackbar.message}
        </Alert>
        </Snackbar>
    </>
    );
}