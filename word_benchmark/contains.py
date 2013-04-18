from optparse import OptionParser
import multiprocessing
import time


def contains(needle, haystack, results):
    out = []
    for word in haystack:
        if needle in word:
            out.append(word)
    results.put(len(out))


def get_haystack(path):
    with open(path) as f:
        return f.read().splitlines()


if __name__ == '__main__':
    parser = OptionParser()
    parser.add_option('-w', '--word', dest='word')
    parser.add_option('-p', '--path', dest='path')
    parser.add_option('-k', '--workers', dest='workers', type="int")
    (options, args) = parser.parse_args()

    num_cpus = multiprocessing.cpu_count()
    if not options.workers:
        options.workers = num_cpus
    haystack = get_haystack(options.path)
    size = len(haystack) / options.workers

    start = time.time()
    partitions = zip(*[iter(haystack)] * size)

    results = multiprocessing.Queue()
    jobs = []
    for partition in partitions:
        p = multiprocessing.Process(target=contains, args=(options.word, partition, results))
        jobs.append(p)
        p.start()

    total = 0
    for job in jobs:
        job.join()
        total += results.get()
    print "Matched: {0}".format(total)
    print "Critical section: {0:f} seconds".format(time.time() - start)
