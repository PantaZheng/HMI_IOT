<template>
<div>
  <assistant  v-if="role == 'assistant'"></assistant>
  <teacher v-else-if="role == 'teacher'"></teacher>
  <others v-else></others>
</div>
</template>

<script>
import Assistant from './assistant/assistant'
import Teacher from './teacher/teacher'
import Others from './others/others'

export default {
  name: 'newProject',
  components: {
    Assistant,
    Teacher,
    Others
  },
  data () {
    return {
      role: 'teacher'
    }
  }
}
</script>

<style scoped>

</style>
