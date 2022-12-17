<template>
    <main>
        <MainHeader />
        <div class="box">
            <div class="active" v-for="d, i in this.$store.state.boxes" :key="i">
                <ActiveBox :box="d" />
            </div>
            <NewBoxVue />
        </div>
    </main>
</template>

<script>
import ActiveBox from '@/components/active box/ActiveBox.vue';
import MainHeader from '@/components/MainHeader.vue';
import NewBoxVue from '@/components/new box/NewBox.vue';
export default {

    components: { NewBoxVue, MainHeader, ActiveBox }
}
</script>

<style scoped>
main {
    max-width: 100vw;
    height: 100%;
    position: fixed;
    z-index: 9;
    
}

.box {
    /* background-color: antiquewhite; */
    overflow-x: auto;
    display: flex;
    height: 80%;
}

h1 {
    margin: 0;
}
</style>