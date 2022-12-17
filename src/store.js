import { createStore } from "vuex";

var store = createStore({
  state() {
    return {
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
    newBox(c, payload) {
      c.boxes.push(payload);
    },
    newTask(c, payload) {
      c.tasks.unshift(payload);
    },
  },
});

export default store;
