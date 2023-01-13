<script lang="ts" setup>
import {computed, onMounted, reactive, watch} from "vue";
import {GetRepositoryInfo} from "../../../wailsjs/go/main/Api";
import {saveSelectedProject} from "../../utils/storage";

const emits = defineEmits(["close"])
const props = defineProps(['project'])
const data = reactive({
  repository: undefined,
  reposForEdit: ""
})

function closeProject() {
  emits("close")
}

function reposInput($event: any) {
  data.reposForEdit = $event.target.value
}

function saveProjectConfiguration() {
  const finalRepos = []
  let orgRepos = data.reposForEdit.split("\n")
  for(let i=0;i<orgRepos.length;i++){
    if(orgRepos[i].trim()!==""){
      finalRepos.push(orgRepos[i])
    }
  }
  props.project.organization_packages = finalRepos
  saveSelectedProject(props.project)
}

onMounted(() => {
  data.reposForEdit = props.project.organization_packages?.join("\n")
  GetRepositoryInfo(props.project.path).then((result: any) => {
    data.repository = result
  })
})
</script>

<template>
  <div class="card">
    <div class="card-header">
      <h3 class="card-title">Project Summary</h3>
      <div class="text-right">
        <div class="btn btn-sm btn-default action spaced" title="Project Configuration" data-toggle="modal" data-target="#modal-config-project" >
          <i class="fa fa-cogs" />
        </div>
        <div class="btn btn-sm btn-danger action" title="Close Project" v-on:click="closeProject">
          <i class="fa fa-close" />
        </div>
      </div>
    </div>
    <div class="card-body">
      <div class="row">
        <div class="col-12 col-md-4 col-lg-4">
          <div class="text-muted">
            <p class="text-sm">
              <b class="d-block">Project Name</b>
              {{ props.project.name }}
            </p>
          </div>
        </div>
        <div class="col-12 col-md-4 col-lg-4">
          <div class="text-muted">
            <p class="text-sm">
              <b class="d-block">Main Package</b>
              {{ props.project.package }}
            </p>
          </div>
        </div>
        <div class="col-12 col-md-4 col-lg-4">
          <div class="text-muted">
            <p class="text-sm">
              <b class="d-block">Project Path</b>
              {{ props.project.path }}
            </p>
          </div>
        </div>
        <div class="col-12 col-md-4 col-lg-4">
          <div class="text-muted">
            <p class="text-sm">
              <b class="d-block">Repository URL</b>
              {{ (data.repository==undefined) ? "-":data.repository["url"] }}
            </p>
          </div>
        </div>
        <div class="col-12 col-md-4 col-lg-4">
          <div class="text-muted">
            <p class="text-sm">
              <b class="d-block">Current Branch</b>
              {{ (data.repository==undefined) ? "-":data.repository["branch"] }}
            </p>
          </div>
        </div>
        <div class="col-12 col-md-4 col-lg-4">
          <div class="text-muted">
            <p class="text-sm">
              <b class="d-block">Current Revision</b>
              {{ (data.repository==undefined) ? "-":data.repository["revision"] }}
            </p>
          </div>
        </div>
      </div>
    </div>
  </div>

  <div class="modal fade" id="modal-config-project" style="display: none;" aria-hidden="true">
    <div class="modal-dialog modal-lg">
      <div class="modal-content">
        <div class="modal-header">
          <h4 class="modal-title">Go Project Configuration</h4>
          <button type="button" class="close" data-dismiss="modal" aria-label="Close">
            <span aria-hidden="true">×</span>
          </button>
        </div>
        <div class="modal-body">
          <div>
            <div class="form-group">
              <label>Same Organization Repositories</label>
              <textarea class="form-control" rows="5"
                        placeholder="Each line must refer to only one repository"
                        v-bind:value="data.reposForEdit" v-on:input="reposInput"></textarea>
            </div>
          </div>
        </div>
        <div class="modal-footer justify-content-between">
          <button type="button" class="btn btn-default" data-dismiss="modal">Cancel</button>
          <button type="button" class="btn btn-primary" data-dismiss="modal" v-on:click="saveProjectConfiguration">Save</button>
        </div>
      </div>

    </div>
  </div>
</template>

<style scoped>
.action {
  width: 40px;
}
.spaced {
  margin-right: 10px;
}
</style>
