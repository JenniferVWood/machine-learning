package org.buskersguidetotheuniverse.mockweatherapi.controller;

import org.springframework.stereotype.Controller;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.ResponseBody;

import java.util.HashMap;
import java.util.Map;

@Controller
@RequestMapping("stations/local/observations")
public class ObservationsController {

    /**
     * Support dev-testing of the code that reads the weather API
     * @return test data
     */
    @GetMapping("/current")
    public @ResponseBody Map<String,Object> currentObservations() {
        Map<String, Object> ret = new HashMap<>();
        ret.put("test", "test");
        return ret;
    }
}
