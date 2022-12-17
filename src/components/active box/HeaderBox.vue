<template>
        <div class="header">
            <h2>{{ box.name }}</h2>
            <p>i</p>
        </div>
</template>

<script>

export default {
    props:["box"]
}
</script>

<style scoped>
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
p {
    cursor: pointer;
    width: 20px;
    height: 20px;
    border: solid rgb(0, 41, 176) 1px;
    border-radius: 50%;
    margin-right: 0.5rem;
}
</style>