<template>
  <div>
    <v-item v-slot:default="{ active, toggle }">
      <v-card
        class="d-flex align-center"
        :color="color"
        @click="toggle"
        height="60"
      >
        <v-scroll-y-transition>
          <div
            class="display-100 flex-grow-1 text-center"
          >
            <div style="color: white;" class="font-weight-bold">
              {{ message }}
              {{ num }}
            </div>
          </div>
        </v-scroll-y-transition>
      </v-card>
    </v-item>
  </div>
</template>


<script lang="ts">
import { Vue, Component, Prop } from 'vue-property-decorator';

@Component({
  componens: {
  },
})
export default class SquarePanel extends Vue {
  @Prop()
  private num: number;
  @Prop()
  private message: string;

  private color: string = 'blue-grey';

  private created(): void {
    if(this.message == '解決済') this.color = '#5cb85c';
    else if(this.message == '未解決') this.color = '#f0ad4e';
  }
}
</script>

<style>

</style>
