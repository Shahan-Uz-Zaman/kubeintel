import axios from "axios";

const API = "http://localhost:8080/api";

export const getEvents = async (namespace = "default") => {
    const response = await axios.get(`${API}/events`, {
        params: { namespace }
    });

    return response.data.events;
};