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
import SupplierInfoCard from "@//components/pages/smart-filter/SupplierInfoCard.vue";

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
          title: "",
          search: "",
          type: "",
        },
        {
          id: 1,
          title: "",
          search: "",
          type: "",
        },
        {
          id: 2,
          title: "",
          search: "",
          type: "",
        },
      ],
      filterableFields: [
        {
          label: "ОГРН",
          value: "ogrn",
        },
        {
          label: "Название",
          value: "name",
        },
        {
          label: "КПП",
          value: "kpp",
        },
        {
          label: "Выполненные тендеры",
          value: "successful_tenders",
        },
        {
          label: "Невыполненные тендеры",
          value: "unsuccessful_tenders",
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
  methods: {
    addFilter() {
      this.filters.push({
        id: this.filters.length,
        type: "",
        title: "",
        search: "",
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