<template>
    <input type="text" v-model="name" v-on:keyup.enter="save" placeholder="New Task">
</template>

<script>
export default {
    data() {
        return {
            name: ""
        }
    },
    methods: {
        save() {
            if (this.name == "") {
                return this.inpErr = true
            }
            console.log(this.name)
            this.$store.dispatch("newTask", {
                name: this.name,
                status: "Requested"
            })
            this.name = ""
        }
    }
}
</script>

<style scoped>
input {
    background-color: rgba(255, 255, 255, 0.541);
    border: solid 1px rgb(169, 169, 169);
    font-size: 1.2rem;
    padding: 0.5rem 1rem;
    border-radius: 8px;
    width: 80%;
    margin: 0.3rem 1rem;

}

.inpErr {
    border-color: red;
}
</style>