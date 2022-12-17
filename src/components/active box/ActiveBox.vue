<template>
    <div class="box">
        <HeaderBox :box="box"/>
        <div class="bodies">
            <NewTask v-if="box.name == 'Requested'"/>
            <div class="tasks" v-for="d, i in this.$store.state.tasks" :key="i">
                <ActiveTask v-if="d.Status == box.name" :task="d"/>
            </div>
        </div>
    </div>
</template>

<script>
import ActiveTask from './ActiveTask.vue';
import HeaderBox from './HeaderBox.vue';
import NewTask from './NewTask.vue';


export default {
    props: ["box", "bindex"],
    methods:{
       
    },
    components: { HeaderBox, NewTask, ActiveTask }
}
</script>

<style scoped>
.box {
    width: 300px;
    min-height: 80%;
    background-color: rgba(255, 255, 255, 0.948);
    border-radius: 8px;
    box-shadow: 0 0 3px 3px rgba(255, 255, 255, 0.05);
    margin: 1rem;
    display: flex;
    flex-direction: column;
}
.header {
    width: 100%;
    display: flex;
    align-items: center;
    justify-content: space-between;
    border-bottom: solid rgb(0, 41, 176) 1px;
    height: min-content;
    /* padding: 0 1rem; */
}
h2 {
    margin: 0.5rem;
    text-align: center;
}
.bodies {
    width:100%;
    display: flex;
    flex-direction: column;
    align-items: center;
}
.tasks {
    width:100%;
}

</style>