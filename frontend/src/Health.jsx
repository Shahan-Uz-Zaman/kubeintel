const [health,setHealth]=useState(null)

useEffect(()=>{

axios.get("/api/health")

.then(res=>setHealth(res.data))

},[])