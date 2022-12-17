import axios from "axios";
import { createStore } from "vuex";

var store = createStore({
  state() {
    return {
      loading:{
        task:false,
      },
      backend:{
        url:"http://127.0.0.1:8090"
      },
      boxes: [
        {
          name: "Requested",
          tag: "main",
          description: "-",
          color: "blue",
          priority: "normal",
        },
        {
          name: "Progress",
          tag: "main",
          description: "-",
          color: "blue",
          priority: "normal",
        },
        {
          name: "Done",
          tag: "main",
          description: "-",
          color: "blue",
          priority: "normal",
        },
      ],
      box: {
        name: "",
        tag: "",
        description: "",
        color: "",
        priority: "",
      },
      tasks: [
        {
            name: "Fixing Bugs on all pages",
            status: "Requested",
          },
          {
            name: "Nexxxxxttt",
            status: "Progress",
          },
          {
            name: "Help Page Starter",
            status: "Done",
          },
      ],
      task: {
        name: "",
        status: "requested",
      },
    };
  },
  mutations: {
    GET_TASKS(c, payload){
      c.tasks = payload
    },
    NEW_TASK(c, payload) {
      c.tasks.unshift(payload);
    },
    NEW_BOX(c, payload) {
      c.boxes.push(payload);
    },
  },
  actions: {
    async getTasks(c){
      c.state.loading.task = true
      try {
        let result = await axios.get(
          c.state.backend.url+"/task/tasks",
          )
        if (result.status == 200){
          c.commit("GET_TASKS", result.data.data.Value)
          c.state.loading.task = false
        }
      } catch (e) {
        alert(e)
      }
    },
    async newTask(c, payload) {
      try {
        let result = await axios.post(
          c.state.backend.url+"/task/store",
          payload,
          {
            headers:{
              "Content-Type":"application/json"
            }
          }
          )
          console.log(result)
          if (result.status == 200){
            c.commit("NEW_TASK", payload)
          }
      } catch (e) {
        alert(e)
      }
    }
  }
});

export default store;
