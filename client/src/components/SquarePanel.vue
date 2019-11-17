<template>
  <div>
    <!-- before -->
    <!--v-item v-slot:default="{ active, toggle }">
      <v-list-item-avatar
        class="tile"
        size="60"
        :color="color"
        tile
      >
        <v-scroll-y-transition>
          <div
            class="display-100 flex-grow-1 text-center"
          >
            <div style="color: white;" class="font-weight-bold">
              {{ message }}
              <br>
              {{ num }}
            </div>
          </div>
        </v-scroll-y-transition>
      </v-list-item-avatar>
    </v-item-->
    
    <!-- FIXME  最悪の分岐-->
    <v-chip
      :color="color"
      text-color="white"
      
    >
      {{ message }}
    </v-chip>

  </div>
</template>


<script lang="ts">
import { Vue, Component, Prop } from 'vue-property-decorator';

@Component
export default class SquarePanel extends Vue {
  @Prop()
  private num!: number;
  @Prop()
  private message!: string;
  @Prop()
  private color!: string;
}
</script>

<style>
.tile{
  border-radius: 15%;
}
</style>
