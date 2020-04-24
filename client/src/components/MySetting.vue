<template>
  <div class="myprofile">
    <UserUpdateForms />
    <WithdrawButton />
  </div>
</template>

<script lang="ts">
import { Component, Vue, Watch } from 'vue-property-decorator';
import WithdrawButton from '@/components/WithdrawButton.vue';
import UserUpdateForms from '@/components/UserUpdateForms.vue';

@Component({
  components: {
    WithdrawButton,
    UserUpdateForms,
  },
})

export default class MySetting extends Vue {
}
</script>
