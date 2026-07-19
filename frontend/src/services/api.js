import axios from "axios";

const api = axios.create({
    baseURL: "http://localhost:8080/api",
    headers: {
        "Content-Type": "application/json",
    },
});

export const getDashboard = () => api.get("/dashboard");
export const getDeployments = async () => {

    return await api.get("/deployments");

};
export const createDeployment = async (data) => {

    return await api.post("/deployments", data);

};
export const deleteDeployment = async (name) => {

    return await api.delete(`/deployments/${name}`);

};
export const restartDeployment = async (name) => {

    return await api.post(`/deployments/${name}/restart`);

};
export const scaleDeployment = async (name, replicas) => {

    return await api.put(

        `/deployments/${name}/scale`,

        {

            namespace: "default",

            replicas,

        }

    );

};
export default api;