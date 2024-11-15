# app structure

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

## templates
- name
- repository_app
- building
   {
      on: { event: GIT_PUSH, data: { branch: develop } },
      execute: { 
         type: PIPELINE, 
         data: { 
            pipeline_id: '1', 
            input: {
         service_name: 'devex-api',
                  service_repository: 'https://github.com/raphaelbh/devex-api'
            }  
         }
      }
   }

## secrets
- key
- value

## services
- name
- repository_app
- template_id
- env_variables (para cada ambiente cadastrado)
   { environment_id: {}, environment_id: {} }

## service_events
- application_id
- type: BUILD | DEPLOY
- trigged_at
- trigged_by
- data: {}

----------

## pipelines 
- name
- definition
   {
      "commands": [
         "echo go service building..."
      ]
   }

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