Feature: List created instances
    So I can see what has been provisioned
    As a cloud engineer
    I want to list the instances using filters

Scenario: The instances are requested without using a filter
    Given the following instances
        | 12354433 | ubuntu-latest | 1Gb | standard-1 | haproxy    |
        | 34534534 | ubuntu-latest | 1Gb | standard-2 | app1       |
        | 76867678 | postgres-20   | 1Gb | gpu-1      | prd-sth-au |
    When the list of instances is requested
    Then the response will be
        """
        {
            "instances": [
                {
                    "id": "12354433",
                    "image": "ubuntu-latest",
                    "swapSize": "1Gb",
                    "type": "standard-1",
                    "label": "haproxy"
                },
                {
                    "id": "34534534",
                    "image": "ubuntu-latest",
                    "swapSize": "1Gb",
                    "type": "standard-2",
                    "label": "app1"
                },
                {
                    "id": "76867678",
                    "image": "postgres-20",
                    "swapSize": "1Gb",
                    "type": "gpu-1",
                    "label": "prd-sth-au"
                }
            ]
        }
        
        """
