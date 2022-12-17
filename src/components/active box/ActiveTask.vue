<template>
    <div class="task"
        :class="[task.status == 'Requested' ? 'requested' : '' || task.status == 'Progress' ? 'progress' : 'done']"
    >
        <div
        @click="backB(task)"
            v-if="task.status != 'Requested'" class="b back">{{ back }}</div>
        <p :class="[task.status == 'Requested' ? 'left' : '']">{{ task.name }}</p>
        <div 
        v-if="task.status != 'Done'"
        class="b next" @click="nextB(task)"> {{ next }}</div>
    </div>
</template>

<script>
export default {
    props: ['task'],
    data(){
        return{
            next:">",
            back:"<"
        }
    },
    methods:{
        nextB(d){
            d.status == "Requested" ? d.status = "Progress"
            : d.status = "Done"
        },
        backB(d){
            d.status == "Done" ? d.status = "Progress"
            : d.status = "Requested"
        }
    }

}
</script>

<style scoped>
.task {
    display: flex;
    align-items: center;
    justify-content: space-between;
    background-color: rgba(255, 255, 255, 0.541);
    border: solid 1px  rgb(0, 41, 176);
    font-size: 1.2rem;
    border-radius: 8px;
    width: 90%;
    margin: 0.3rem 0.8rem;
    text-align: start;
    font-size: 0.95rem;
}
.b {
    padding: 0.5rem 1rem;
    height: 100%;
    color: rgb(0, 43, 161);
    font-weight: bold;
    cursor: pointer;
}
.back {
    /* background-color: orange; */
    border-radius: 8px 0 0 8px;
}
.next {
    /* background-color: rgb(0, 41, 176); */
    border-radius: 0 8px 8px 0;
}
.requested {
    justify-content: space-between;
}
.progress {
    justify-content: space-between;
}
.done {
    justify-content: flex-start;
}
.left {
    
    padding-left: 1rem;
}
</style>