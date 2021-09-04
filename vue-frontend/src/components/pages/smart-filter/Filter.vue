<template>
  <div class="filter">
    <div class="filter-top-row">
      <el-select
        :value="title"
        filterable
        placeholder="Фильтровать по..."
        @change="onTitleChange"
      >
        <el-option
          v-for="option in filterableFields"
          :key="option.value + '-' + id"
          :label="option.label"
          :value="option.value"
        >
        </el-option>
      </el-select>
      <el-select
        :value="type"
        filterable
        placeholder="тип фильтрации..."
        @change="onFilterTypeChange"
      >
        <el-option
          v-for="option in filterTypes"
          :key="option.value + '-' + id"
          :label="option.label"
          :value="option.value"
        >
        </el-option>
      </el-select>
      <el-button @click="removeFilter" class="delete-button">
        <i class="el-icon-delete delete-icon"></i>
      </el-button>
    </div>
    <el-input :placeholder="title" :value="search" @input="onInput"></el-input>
  </div>
</template>

<script>
import _ from "lodash";

export default {
  name: "FilterComponent",
  props: {
    id: {
      required: true,
      type: Number,
    },
    type: {
      required: true,
      type: String,
    },
    title: {
      required: true,
      type: String,
    },
    search: {
      required: true,
      type: String,
    },
    filterableFields: {
      required: true,
      type: Array,
    },
  },
  data() {
    return {
      filterTypes: [
        {
          label: "Отсечение снизу",
          value: "low_cut",
        },
        {
          label: "Отсечение сверху",
          value: "high_cut",
        },
        {
          label: "равенство",
          value: "eq",
        },
        {
          label: "вхождение",
          value: "like",
        },
        {
          label: "вхождение (без регистра)",
          value: "ilike",
        },
      ],
    };
  },
  methods: {
    onInput: _.debounce(function (val) {
      this.$emit("update:search", val);
    }),
    onFilterTypeChange(val) {
      this.$emit("update:type", val);
    },
    onTitleChange(val) {
      this.$emit("update:title", val);
    },
    removeFilter() {
      console.log("to delete", this.id);
      this.$emit("delete", this.id);
    },
  },
};
</script>

<style scoped>
.filter {
  display: flex;
  flex-direction: column;
  margin-right: 32px;
  align-items: flex-start;
  position: relative;
}

.filter-top-row {
  display: flex;
}

.filter-label {
}

.delete-button {
}

.delete-icon:hover {
  color: red;
  cursor: pointer;
}
</style>