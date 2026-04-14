export namespace service {
	
	export class NetworkInfo {
	    ip: string;
	    subnet: string;
	    gateway: string;
	
	    static createFrom(source: any = {}) {
	        return new NetworkInfo(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.ip = source["ip"];
	        this.subnet = source["subnet"];
	        this.gateway = source["gateway"];
	    }
	}

}

