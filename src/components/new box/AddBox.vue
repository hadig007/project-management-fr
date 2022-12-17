<template>
    <main>
        <input :class="[inpErr == true ? 'inpErr' : '']" v-on:keyup.enter="save" type="text" v-model="name" placeholder="New box name">
        <button @click="save">Save</button>
    </main>
</template>

<script>
export default {
    data(){
        return{
            name:"",
            inpErr: false
        }
    },
    methods:{
        save(){
            if(this.name == ""){
                return this.inpErr = true

            }
            console.log(this.name)
            this.$store.dispatch("newBox", {
                name:this.name,
                tag:"-",
                description:"-",
                color:"-",
                priority:"-"
            })
            this.name = ""
        }
    }
}
</script>

<style scoped>
main{
    width: 100%;
    height: 100%;
    display: flex;
    align-items: center;
    justify-content: center;
    flex-direction: column;
}
p {
    font-size: 5rem;
}
input {
    background-color: rgba(255, 255, 255, 0.454);
    border: solid 1px rgb(255, 255, 255);
    font-size: 1.2rem;
    padding: 0.5rem 1rem;
    border-radius: 8px;
}
.inpErr {
    border-color: red;
}
button {
    padding: 0.5rem 1rem;
    background-color: rgb(0, 41, 176);
    border: solid 1px transparent;
    border-radius: 8px;
    margin: 0.5rem;
    color: white;
}
</style>