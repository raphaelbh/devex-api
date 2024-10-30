# app structure

## templates
- name
- repository_app
- building
   {
      on: { event: GIT_PUSH, data: { branch: develop } },
      execute: { 
         type: PIPELINE, 
         data: { 
            pipeline_id: 1, 
            input: {
               service_name: vira da consulta,
               service_repository: vira da consulta
               tag: calculado a partir da consulta
            }  
         }
      }
   }

## applications
- name
- repository_app
- template_id
- env_variables (para cada ambiente cadastrado)
   { environment_id: {}, environment_id: {} }

## application_variables
- application_id
- type: VARIABLE | SECRET
- key
- value

## application_events
- application_id
- type: BUILD | DEPLOY
- trigged_at
- trigged_by
- data: {}

## environments
- name
- deployment
{
   on: { event: PIPELINE_EXECUTION_COMPLETED, data: { pipeline_id: 1 } },
   execute: { 
      type: PIPELINE, 
      data: { 
         pipeline_id: 2, 
         input: { 
            service_name: vira do evento
            tag: vira do evento
         }
      }
   }
}

----------

## pipelines 
- name
- version: 1
- definition: { ... }

## pipeline_executions
- pipeline_id
- status [ WAITING_EXECUTION | IN_PROGRESS | CANCELED | FAIL | COMPLETED ]
- input
- logs
- elapsed_time
- starting_at
- finishing_at
- trigged_at
- trigged_by
- trigger_type [ MANUAL | PIPELINE | GIT | API ]

## pipeline_variables
- pipeline_id
- type: VARIABLE | SECRET
- key
- value