<template>
    <header>
        <h1>{{ title }}</h1>
        <p>{{ cTime }}</p>
    </header>
</template>

<script>
// import moment from "moment"
export default {
    created() {
        setInterval(() => {
            this.updateTime()
        }, 1000)
    },
    methods: {
        updateTime() {
            var t = new Date()
            var time = `${t.getHours()}:${t.getMinutes()}:${t.getSeconds()}`
            var day = this.getDayString(t.getDay())
            var date = t.getDate()
            var moon = this.getMoonString(t.toDateString())
            var year = t.getFullYear()
            this.cTime = `${day}, ${date} ${moon} ${year} ${time}`
            return this.cTime
        },
        getDayString(d) {
            if (d == 1) {
                return d = "Senin"
            }
            else if (d == 2){
                return d = "Selasa"
            }
            else if (d == 3){
                return d = "Rabu"
            }
            else if (d == 4){
                return d = "Kamis"
            }
            else if (d == 5){
                return d = "Jumat"
            }
            else if (d == 6){
                return d = "Sabtu"
            }
            else if (d == 7){
                return d = "Minngu"
            }else {
                d = "err"
            }
        }, 
        getMoonString(d) {
            d = d.split(" ")[1]
            if (d == "Jan") {
                return d = "Januari"
            }
            else if (d == "Feb"){
                return d = "February"
            }
            else if (d == "Mar"){
                return d = "Maret"
            }
            else if (d == "Apr"){
                return d = "April"
            }
            else if (d == "May"){
                return d = "Mei"
            }
            else if (d == "Jun"){
                return d = "Juni"
            }
            else if (d == "Jul"){
                return d = "Juli"
            }
            else if (d == "Aug"){
                return d = "Agustus"
            }
            else if (d == "Sep"){
                return d = "September"
            }
            else if (d == "Oct"){
                return d = "Oktober"
            }
            else if (d == "Nov"){
                return d = "November"
            }
            else if (d == "Dec"){
                return d = "Desember"
            }
            else {
                d = "err"
            }
        }
    },
    data() {
        return {
            title: "Project Managament",
            cTime: "loading .."
        }
    }
}
</script>


<style scoped>
header {
    border-bottom: solid white 1px;
}
h1 {
    color: white;
    margin: 0;
    padding: 0.5rem;
    padding-bottom: 0;
}
p {
    color: white;
    margin: 0;
}
</style>