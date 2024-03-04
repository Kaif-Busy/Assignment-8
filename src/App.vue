<template>
  <div id="app">
    <table>
      <thead>
        <tr>
          <th>Partner Users</th>
          <th>Active Lead</th>
          <th>Unassigned Leads</th>
          <th>Active Prospect</th>
          <th>Active Opportunity</th>
        </tr>
      </thead>
      <tbody>
        <tr v-for="row in tableData" :key="row.id">
          <td>{{ row.PartnerUsers }}</td>
          <td>{{ row.ActiveLead }}</td>
          <td>{{ row.UnassignedLeads }}</td>
          <td>{{ row.ActiveProspect }}</td>
          <td>{{ row.ActiveOpportunity }}</td>
        </tr>
      </tbody>
    </table>
  </div>
</template>

<script>
import axios from 'axios';

export default {
  data() {
    return {
      tableData: [],
    };
  },
  mounted() {
    this.fetchData();
  },
  methods: {
    async fetchData() {
      try {
        const response = await axios.get('http://localhost:8080/crm');
        this.tableData = response.data;
      } catch (error) {
        console.error('Error fetching data:', error);
      }
    },
  },
};
</script>

<style>
table {
  border-collapse: collapse;
  width: 100%;
}

th, td {
  border: 1px solid #dddddd;
  text-align: left;
  padding: 8px;
}

th {
  background-color: #f2f2f2;
}
</style>
