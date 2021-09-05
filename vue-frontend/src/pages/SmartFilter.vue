<template>
  <div class="center">
    <h2>Поиск поставщиков</h2>
    <div class="filters-container">
      <FilterComponent
        v-for="filter in filters"
        :key="filter.id"
        :id="filter.id"
        :type.sync="filter.type"
        :title.sync="filter.title"
        :search.sync="filter.search"
        :filterableFields="filterableFields"
        @delete="removeFilter"
      />
      <el-button type="primary" @click="addFilter">
        <i class="el-icon-plus"></i>
      </el-button>
    </div>
    <div class="suppliers-container">
      <SupplierInfoCard
        v-for="(supplier, ind) in suppliers"
        :key="supplier.ogrn"
        :supplier="supplier"
        :ind="ind"
      />
    </div>
  </div>
</template>

<script>
import FilterComponent from "@/components/pages/smart-filter/Filter.vue";
import SupplierInfoCard from "@/components/pages/smart-filter/SupplierInfoCard.vue";

import _ from "lodash";

export default {
  name: "SmartFilterPage",
  components: {
    FilterComponent,
    SupplierInfoCard,
  },
  data() {
    return {
      suppliers: [],
      filters: [
        {
          id: 0,
          type: "",
          title: "",
          search: "",
          searchType: "",
        },
        {
          id: 1,
          type: "",
          title: "",
          search: "",
          searchType: "",
        },
        {
          id: 2,
          type: "",
          title: "",
          search: "",
          searchType: "",
        },
      ],
      filterableFields: [
        {
          label: "ОГРН",
          value: "ogrn",
          searchType: "number",
        },
        {
          label: "Название",
          value: "name",
          searchType: "string",
        },
        {
          label: "КПП",
          value: "kpp",
          searchType: "number",
        },
        {
          label: "Выполненные тендеры",
          value: "successful_tenders",
          searchType: "number",
        },
        {
          label: "Невыполненные тендеры",
          value: "unsuccessful_tenders",
          searchType: "number",
        },
      ],
    };
  },
  created() {
    this.$http.hasura
      .post("/", {
        query: require("@/gql/qFetchSome.js").default,
      })
      .then((resp) => {
        if (resp.data.errors) {
          return;
        }
        const blacklist = resp.data.data.supplier_blacklist;

        this.suppliers = resp.data.data.suppliers.map((s) => ({
          ...s,
          blacklisted: !!blacklist.find((bs) => bs.supplier_inn === s.inn),
        }));
      });
  },
  watch: {
    filters: {
      deep: true,
      handler() {

        this.applyFilters();
      },
    },
  },
  methods: {
    applyFilters: _.debounce(function () {
      let filterObjects = this.filters
        .filter((f) => !!f.search && !!f.type && !!f.title)
        .map((f) => {
          let res = {};
          res[f.title] = {};

          let search;
          if (f.type === "_like" || f.type === "_ilike") {
            search = `"%${f.search}%"`;
          } else {
            search = f.search;
          }

          res[f.title][f.type] = search;

          return res;
        });
      console.log("filterObjects", filterObjects)
      console.log(require("@/gql/qSearchSuppliers.js").default(filterObjects));


      // this.$http.hasura
      //   .post("/", {
      //     query: require("@/gql/qSearchSuppliers.js").default({
      //       filterObjects,
      //     }),
      //   })
      //   .then((resp) => {
      //     if (resp.data.errors) {
      //       return;
      //     }
      //     this.suppliers = this.applyFilters(resp.data.data.suppliers);
      //   });


    }, 300),
    addFilter() {
      this.filters.push({
        id: this.filters.length,
        type: "",
        title: "",
        search: "",
        searchType: "",
      });
    },
    removeFilter(fid) {
      this.filters = this.filters.filter((f) => f.id !== fid);
      this.filters.forEach((f, ind) => {
        f.id = ind;
      });
    },
  },
};
</script>

<style scoped>
.center {
  display: block;
  margin-left: auto;
  margin-right: auto;
}

.filters-container {
  border: 1px solid gray;
  display: flex;
  align-items: center;
  width: 80%;
  margin-left: auto;
  margin-right: auto;
  border-radius: 3px;
  padding-top: 32px;
  padding-bottom: 32px;
  padding-left: 32px;
  padding-right: 32px;
}

.suppliers-container {
  margin-top: 20px;
  display: flex;
  flex-direction: column;
  align-items: stretch;
  width: 80%;
  margin-left: auto;
  margin-right: auto;
  border-radius: 3px;
  padding-top: 32px;
  padding-bottom: 32px;
  padding-left: 32px;
  padding-right: 32px;
}
</style>