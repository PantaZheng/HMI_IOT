<template>
  <div>
    <assistant-header></assistant-header>
    <assistant-form></assistant-form>
  </div>
</template>

<script>
import AssistantHeader from './components/header'
import AssistantForm from './components/form'

export default {
  name: 'assistant',
  components: {
    AssistantHeader,
    AssistantForm
  }
}
</script>

<style scoped>

</style>
