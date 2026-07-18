import axios from "axios";

const API = "http://localhost:8080/api/monitoring";

export const getClusterMetrics = () =>
  axios.get(`${API}/cluster`);

export const getNodeMetrics = () =>
  axios.get(`${API}/nodes`);

export const getPodMetrics = () =>
  axios.get(`${API}/pods`);

export const getStorage = () =>
    axios.get(`${API}/storage`);

export const getNetwork = () =>
    axios.get(`${API}/network`);