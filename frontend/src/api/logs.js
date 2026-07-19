import axios from "axios";

const API = "http://localhost:8080/api";

export const getNamespaces = async () => {
    const res = await axios.get(`${API}/namespaces`);
    return res.data.namespaces; // or res.data.data, depending on your API
};

export const getPods = async (namespace = "default") => {
    const res = await axios.get(`${API}/pods`, {
        params: { namespace }
    });
    return res.data.pods; // or res.data.data
};

export const getLogs = async (namespace, pod) => {
    const res = await axios.get(`${API}/logs`, {
        params: { namespace, pod }
    });

    return res.data.logs;
};